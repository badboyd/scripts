package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

	protov1 "github.com/badboyd/scripts/grpcserver/proto/v1"
)

type Server struct{}

func (s *Server) ListSomething(ctx context.Context, req *protov1.ListSomethingRequest) (*protov1.ListSomethingResponse, error) {
	log.Printf("Request limit %d, next %d", req.Limit, req.Next)

	res := &protov1.ListSomethingResponse{
		Foofoo: make([]*protov1.ListSomethingResponse_Data, 0, req.Limit),
	}

	for i := 0; i < int(req.Limit); i++ {
		res.Foofoo = append(res.Foofoo, &protov1.ListSomethingResponse_Data{
			Id:  int64(rand.Intn(1999)),
			Foo: fmt.Sprint(rand.Intn(100000000)),
		})
	}

	return res, nil
}

func main() {
	var (
		zapLogger, _ = zap.NewProduction()
	)

	logOpts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
		grpc_zap.WithLevels(grpc_zap.DefaultCodeToLevel),
		grpc_zap.WithDecider(func(fullMethodName string, err error) bool {
			// will not log gRPC calls if it was a call to healthcheck and no error was raised
			if err == nil && strings.Contains(fullMethodName, "health") {
				return false
			}

			// by default everything will be logged
			return true
		}),
	}

	// Make sure that log statements internal to gRPC library are logged using the zapLogger as well.
	grpc_zap.ReplaceGrpcLogger(zapLogger)

	gs := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_zap.StreamServerInterceptor(zapLogger, logOpts...),
		),
		grpc.UnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(zapLogger, logOpts...),
		),
	)
	lis, err := net.Listen("tcp", "0.0.0.0:13000")
	if err != nil {
		panic(err)
	}

	reflection.Register(gs)
	protov1.RegisterListSomethingServiceServer(gs, &Server{})

	gs.Serve(lis)
}
