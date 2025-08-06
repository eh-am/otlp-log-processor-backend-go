package main

import (
	"context"
	"log/slog"

	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
)

type dash0LogsServiceServer struct {
	addr string

	collogspb.UnimplementedLogsServiceServer
}

func newServer(addr string) collogspb.LogsServiceServer {
	s := &dash0LogsServiceServer{addr: addr}
	return s
}

func (l *dash0LogsServiceServer) Export(ctx context.Context, request *collogspb.ExportLogsServiceRequest) (*collogspb.ExportLogsServiceResponse, error) {
	slog.DebugContext(ctx, "Received ExportLogsServiceRequest")
	logsReceivedCounter.Add(ctx, 1)

	slog.Info(request.String())
	// Do something with the logs

	// Extract the log
	// TODO Is it a single line? Do we care about multi line?
	// Depending on what it is, parse it
	// Structured logs
	// - If it's JSON, create a custom parser to throw away everything
	//   except the field we care about, to reduce in memory usage
	// Semi structured
	// Unstructured

	// Store the count somewhere

	return &collogspb.ExportLogsServiceResponse{}, nil
}
