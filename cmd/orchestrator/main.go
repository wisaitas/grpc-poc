package main

// import (
// 	"context"
// 	"log"
// 	"net"

// 	"github.com/wisaitas/grpc-poc/pb"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// type orchestratorServer struct {
// 	pb.UnimplementedOrchestratorServiceServer
// 	domainClient pb.DomainServiceClient
// }

// func (s *orchestratorServer) ProcessData(ctx context.Context, req *pb.DataRequest) (*pb.DataResponse, error) {
// 	log.Printf("Orchestrator processing request for ID: %s", req.Id)

// 	// เรียก Domain Service
// 	res, err := s.domainClient.GetData(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// เพิ่ม Logic ของ Orchestrator ต่อท้าย
// 	res.FromService = res.FromService + " -> Orchestrator"
// 	return res, nil
// }

// func main() {
// 	// เชื่อมต่อไปยัง Domain Service
// 	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect to domain: %v", err)
// 	}
// 	defer conn.Close()
// 	domainClient := pb.NewDomainServiceClient(conn)

// 	// เปิด gRPC Server ของตัวเอง
// 	lis, err := net.Listen("tcp", ":50052")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	pb.RegisterOrchestratorServiceServer(s, &orchestratorServer{domainClient: domainClient})

// 	log.Println("Orchestrator Service listening on :50052")
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
