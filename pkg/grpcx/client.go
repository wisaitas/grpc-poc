package grpcx

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCConn[T any] struct {
	Conn   *grpc.ClientConn
	Client T
}

func NewGRPCConn[T any](
	endpoint string,
	newClientFunc func(grpc.ClientConnInterface) T,
) (*GRPCConn[T], error) {
	conn, err := grpc.NewClient(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		return nil, err
	}

	return &GRPCConn[T]{
		Conn:   conn,
		Client: newClientFunc(conn),
	}, nil
}

func (g *GRPCConn[T]) Close() error {
	return g.Conn.Close()
}
