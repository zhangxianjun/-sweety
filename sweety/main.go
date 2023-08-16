package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net/http"
	gw "sweety/proto/business"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// grpc服务地址
	endpoint := "127.0.0.1:50052"
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// HTTP转grpc
	err := gw.RegisterHelloHTTPHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		grpclog.Fatalf("Register handler err:%v\n", err)
	}

	grpclog.Println("HTTP Listen on 8080")
	http.ListenAndServe(":8080", mux)
}
