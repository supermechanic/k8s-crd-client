package server

import (
	"context"
	"fmt"
	"log"
	"model-service/config"
	"model-service/util"
	"net"
	"net/http"

	pb "model-service/api/v1/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

//Run ///
func Run() (err error) {

	conn, err := net.Listen("tcp", config.Config.Endpoint)
	if err != nil {
		fmt.Printf("TCP Listen err:%v\n", err)
	}

	srv := newServer(conn)

	fmt.Printf("gRPC and https listen on: %s\n", config.Config.Endpoint)

	if err = srv.Serve(conn); err != nil {
		log.Printf("ListenAndServe: %v\n", err)
	}

	return err
}

func newServer(conn net.Listener) *http.Server {
	grpcServer := newGrpc()
	gwmux, err := newGateway()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	return &http.Server{
		Addr:    config.Config.Endpoint,
		Handler: util.GrpcHandlerFunc(grpcServer, mux),
	}
}

func newGrpc() *grpc.Server {

	server := grpc.NewServer()

	pb.RegisterModelServiceServer(server, NewModelService())

	return server
}

func newGateway() (http.Handler, error) {
	ctx := context.Background()
	dopts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	if err := pb.RegisterModelServiceHandlerFromEndpoint(ctx, gwmux, config.Config.Endpoint, dopts); err != nil {
		return nil, err
	}

	return gwmux, nil
}
