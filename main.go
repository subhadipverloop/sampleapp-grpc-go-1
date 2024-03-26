package main

import (
	"context"
	"fmt"
	"grpc-sample/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

// implementing simple server for gRPC

type messageServer struct {
	service.UnimplementedMessageServer
}

func (s messageServer) Hello1(ctx context.Context, msg *service.HelloMessageRequest) (*service.HelloMessageResponse, error) {

	// printing client request for hello1 RPC call
	fmt.Println("Received Message from Cliet")
	fmt.Println(msg)

	return &service.HelloMessageResponse{
		Type:       "Response",
		From:       "Hello1",
		Msg:        "Response for RPC hello1",
		ResponseId: 1,
	}, nil
}

func (s messageServer) Hello2(ctx context.Context, msg *service.HelloMessageRequest) (*service.HelloMessageResponse, error) {

	// printing client request for hello2 RPC call
	fmt.Println("Received Message from Cliet")
	fmt.Println(msg)

	return &service.HelloMessageResponse{
		Type:       "Response",
		From:       "Hello2",
		Msg:        "Response for RPC hello2",
		ResponseId: 2,
	}, nil
}

func main() {
	fmt.Println("Inside main.go")

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Cannot create listner %s", err)
	}

	server := grpc.NewServer()

	serv := &messageServer{}

	service.RegisterMessageServer(server, serv)

	fmt.Println("Serving gRPC server on localhost:8080")
	err = server.Serve(lis)

	if err != nil {
		log.Fatalf("can not serve %s", err)
	}

}
