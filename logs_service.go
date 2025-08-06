package main

import (
	"context"
	"log/slog"

	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
)

type dash0LogsServiceServer struct {
	addr string
	svc  Svc

	collogspb.UnimplementedLogsServiceServer
}

type Svc interface {
	Run([]byte)
}

func newServer(addr string, svc Svc) collogspb.LogsServiceServer {
	s := &dash0LogsServiceServer{addr: addr, svc: svc}
	return s
}

func (l *dash0LogsServiceServer) Export(ctx context.Context, request *collogspb.ExportLogsServiceRequest) (*collogspb.ExportLogsServiceResponse, error) {
	slog.DebugContext(ctx, "Received ExportLogsServiceRequest")
	logsReceivedCounter.Add(ctx, 1)

	// TODO:
	// This is wrong, since it only deals with a bare string
	// As opposed to already structured data
	l.svc.Run([]byte(request.String()))

	return &collogspb.ExportLogsServiceResponse{}, nil
}
