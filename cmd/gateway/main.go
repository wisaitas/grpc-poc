package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/wisaitas/grpc-poc/pb"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func main() {
// 	// เชื่อมต่อไปยัง Orchestrator Service
// 	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect to orchestrator: %v", err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewOrchestratorServiceClient(conn)

// 	// สร้าง HTTP Handler
// 	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
// 		id := r.URL.Query().Get("id")
// 		if id == "" {
// 			http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
// 			return
// 		}

// 		// เรียก gRPC ไปยัง Orchestrator
// 		res, err := client.ProcessData(r.Context(), &pb.DataRequest{Id: id})
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		// ตอบกลับเป็น JSON
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(res)
// 	})

// 	log.Println("Gateway HTTP Server listening on :8080")
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		log.Fatal(err)
// 	}
// }
