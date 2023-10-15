// server is a HelloService gRPC server
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	pb "github.com/venim/hello-bazel/proto"
	"github.com/venim/hello-bazel/web/static"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	grpcSrv := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcSrv, new())

	wrappedSrv := grpcweb.WrapServer(grpcSrv, grpcweb.WithCorsForRegisteredEndpointsOnly(false))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if wrappedSrv.IsGrpcWebRequest(r) {
			wrappedSrv.ServeHTTP(w, r)
			return
		}
		http.FileServer(http.FS(static.Static)).ServeHTTP(w, r)
	})
	httpServer := &http.Server{Addr: ":8080", Handler: handler}
	errs := make(chan error, 0)
	go func() {
		lis, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			errs <- fmt.Errorf("net.Listen(): %w", err)
		}
		defer lis.Close()
		if err := httpServer.Serve(lis); err != nil {
			errs <- fmt.Errorf("Serve(): %w", err)
		}
	}()
	go func() {
		lis, err := net.Listen("tcp", "localhost:10123")
		if err != nil {
			errs <- fmt.Errorf("net.Listen(): %w", err)
		}
		defer lis.Close()
		if err := grpcSrv.Serve(lis); err != nil {
			errs <- fmt.Errorf("Serve(): %w", err)
		}
	}()
	defer func() {
		httpServer.Shutdown(ctx)
		grpcSrv.GracefulStop()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	for {
		select {
		case <-stop:
			slog.Warn("ctrl+c captured! shutting down")
			return nil
		case <-ctx.Done():
			return nil
		case err := <-errs:
			return err
		}
	}
}

type svc struct {
	pb.UnimplementedHelloServiceServer
	logger slog.Logger
}

func new() *svc {
	return &svc{
		logger: *slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (svc *svc) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	svc.logger.Info("request recieved",
		slog.String("method", "SayHello"),
		slog.String("name", name),
	)
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	}, nil
}
