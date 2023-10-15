// client is a HelloService gRPC client
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/venim/hello-bazel/proto"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	addr := "localhost:8080"
	cconn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return err
	}
	client := pb.NewHelloServiceClient(cconn)
	res, err := client.SayHello(ctx, &pb.HelloRequest{
		Name: "World",
	})
	if err != nil {
		return err
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("addr", "localhost:8080")
	logger.Info(res.Message)
	return nil
}
