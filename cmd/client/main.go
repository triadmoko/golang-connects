package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	todov1 "golang-connects/gen/proto/todo/v1"
	"golang-connects/gen/proto/todo/v1/todov1connect"
	"log"
	"net"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
)

const tokenHeader = "Acme-Token"

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				// Send a token with client requests.
				req.Header().Set(tokenHeader, "sample")
			} else if req.Header().Get(tokenHeader) == "" {
				// Check token in handlers.
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no token provided"),
				)
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
func newInsecureClient() *http.Client {
	return &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, _ *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr) // see docs for some notes about this :)
			},
		},
	}
}
func main() {
	interceptors := connect.WithInterceptors(NewAuthInterceptor())

	client := todov1connect.NewTodoServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		interceptors,
		connect.WithGRPC(),
	)
	stream := client.List(context.Background())
	sent := 0
	for {
		err := stream.Send(&todov1.ListRequest{Page: 1, Limit: 10})
		if err != nil {
			fmt.Println("sent", sent)
			panic(err)
		}
		sent++
		res, err := stream.Receive()
		if err != nil {
			fmt.Println("err", err)
			return
		}

		prety(res)
		time.Sleep(10 * time.Millisecond)
	}

	// res2, err := client.Read(context.Background(), connect.NewRequest(&todov1.ReadRequest{Id: "1"}))

	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
}
func prety(i any) {
	indent, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(indent))

}
