// server is a HelloService gRPC server
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	pb "github.com/venim/hello-bazel/proto"
	"github.com/venim/hello-bazel/web"
)

func main() {
	if err := run(context.Background(), ":8080"); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, addr string) error {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	svc := new()
	grpcSrv := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcSrv, svc)
	grpc_health_v1.RegisterHealthServer(grpcSrv, health.NewServer())

	wrappedSrv := grpcweb.WrapServer(grpcSrv, grpcweb.WithCorsForRegisteredEndpointsOnly(false))

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(web.Static)))

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcSrv.ServeHTTP(w, r)
			return
		}
		if wrappedSrv.IsGrpcWebRequest(r) {
			wrappedSrv.ServeHTTP(w, r)
			return
		}
		mux.ServeHTTP(w, r)
	})

	httpServer := &http.Server{Addr: addr, Handler: h2c.NewHandler(handler, &http2.Server{})}

	errs := make(chan error, 0)
	go func() {
		slog.Info("server started!", "addr", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil {
			errs <- fmt.Errorf("Serve(): %w", err)
		}
	}()

	defer func() {
		httpServer.Shutdown(ctx)
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
