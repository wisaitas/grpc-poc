package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/wisaitas/grpc-poc/pb"
	"google.golang.org/grpc"
)

type domainServer struct {
	pb.UnimplementedDomainServiceServer
}

func (s *domainServer) GetData(ctx context.Context, req *pb.DataRequest) (*pb.DataResponse, error) {
	log.Printf("Domain received request for ID: %s", req.Id)
	return &pb.DataResponse{
		Result:      fmt.Sprintf("Data for ID %s", req.Id),
		FromService: "Domain",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDomainServiceServer(s, &domainServer{})

	log.Println("Domain Service listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
