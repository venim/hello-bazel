package main

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/venim/hello-bazel/proto"
)

func makeClient(ctx context.Context, addr string) (pb.HelloServiceClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	}
	cconn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, fmt.Errorf("Dial() %w", err)
	}
	client := pb.NewHelloServiceClient(cconn)
	return client, nil
}

func TestGrpcServer(t *testing.T) {
	timeout := 5 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	l, _ := net.Listen("tcp", "")
	addr := l.Addr().String()
	l.Close()
	go run(ctx, addr)

	client, err := makeClient(ctx, addr)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		want string
	}{
		{"World", "Hello, World!"},
		{"bazel", "Hello, bazel!"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := &pb.HelloRequest{
				Name: tc.name,
			}
			res, err := client.SayHello(ctx, req)
			if err != nil {
				t.Fatalf("SayHello(): %v", err)
			}

			want := tc.want
			got := res.GetMessage()

			if diff := cmp.Diff(want, got); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
