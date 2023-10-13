// server is a HelloService gRPC server
package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/venim/hello-bazel/proto"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	lis, err := net.Listen("tcp", "localhost:10123")
	if err != nil {
		return fmt.Errorf("net.Listen(): %w", err)
	}
	defer lis.Close()
	srv := grpc.NewServer()
	pb.RegisterHelloServiceServer(srv, &svc{})
	if err := srv.Serve(lis); err != nil {
		return fmt.Errorf("Serve(): %w", err)
	}
	return nil
}

type svc struct {
	pb.UnimplementedHelloServiceServer
}

func (*svc) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}
