package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"log/slog"
	"net"
	"os"
	"time"

	"dash0.com/otlp-log-processor-backend/internal/manager"
	"dash0.com/otlp-log-processor-backend/internal/parser"
	"github.com/benbjohnson/clock"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	listenAddr            = flag.String("listenAddr", "localhost:4317", "The listen address")
	maxReceiveMessageSize = flag.Int("maxReceiveMessageSize", 16777216, "The max message size in bytes the server can receive")
)

const name = "dash0.com/otlp-log-processor-backend"

var (
	tracer              = otel.Tracer(name)
	meter               = otel.Meter(name)
	logger              = otelslog.NewLogger(name)
	logsReceivedCounter metric.Int64Counter
)

func init() {
	var err error
	logsReceivedCounter, err = meter.Int64Counter("com.dash0.homeexercise.logs.received",
		metric.WithDescription("The number of logs received by otlp-log-processor-backend"),
		metric.WithUnit("{log}"))
	if err != nil {
		panic(err)
	}

	// Use a simpler logger when running locally
	if isDev() {
		debugLvl := new(slog.LevelVar)
		debugLvl.Set(slog.LevelDebug)

		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: debugLvl,
		}))
	} else {
		logger = otelslog.NewLogger(name)
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	slog.SetDefault(logger)
	logger.Info("Starting application")

	// Shut up OTEL when running locally
	if !isDev() {
		// Set up OpenTelemetry.
		otelShutdown, err := setupOTelSDK(context.Background())
		if err != nil {
			return err
		}

		// Handle shutdown properly so nothing leaks.
		defer func() {
			err = errors.Join(err, otelShutdown(context.Background()))
		}()
	}

	flag.Parse()

	slog.Debug("Starting listener", slog.String("listenAddr", *listenAddr))
	listener, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		return err
	}

	// TODO: load config from file
	config := &manager.Config{
		Interval: 30 * time.Second,
		Operations: []parser.ParserConfig{
			{
				Kind:        "regex",
				MatchString: `\{.*\}`,
			},
			{
				Kind:            "json",
				FieldOfInterest: "foo",
			},
		},
	}
	clock := clock.New()

	manager, err := manager.New(config, clock, os.Stdout)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.MaxRecvMsgSize(*maxReceiveMessageSize),
		grpc.Creds(insecure.NewCredentials()),
	)
	collogspb.RegisterLogsServiceServer(grpcServer, newServer(*listenAddr, manager))

	slog.Debug("Starting gRPC server")

	return grpcServer.Serve(listener)
}

func isDev() bool {
	return os.Getenv("DASH0_HOMEEXERCISE_ENV") == "dev"
}
