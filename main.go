package main

import "github.com/kidmodo/qresist/gen/go/sample"
import (
    "context"
    "log"
    "net"
    "google.golang.org/grpc"
)

type server struct {
    sample.UnimplementedSampleServiceServer
}

// GetSample implements the SampleService.GetSample RPC.
func (s *server) GetSample(ctx context.Context, req *sample.SampleMessage) (*sample.SampleMessage, error) {
    log.Printf("Received request: %v", req)
    return &sample.SampleMessage{Name: "Response", Id: req.Id}, nil
}

func main() {
    // Listen on a TCP port
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    // Create a gRPC server
    grpcServer := grpc.NewServer()
    sample.RegisterSampleServiceServer(grpcServer, &server{})

    log.Println("gRPC server is running on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
