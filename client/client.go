package main

import (
	"context"
	"fmt"
	"grpc-sample/service"
	"log"

	"google.golang.org/grpc"
)

// simple client for gRPC

func main() {

	fmt.Println("Inside client.go")

	// var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Can no dial %s", err)
	}
	defer conn.Close()

	c := service.NewMessageClient(conn)

	res, err := c.Hello1(
		context.Background(),
		&service.HelloMessageRequest{
			Type:      "Request",
			From:      "Client go",
			Msg:       "Message from Client GO for Hello1 RPC call",
			RequestId: 100},
	)

	if err != nil {
		log.Fatalf("Can not get details line no- 36 %s", err)
	}
	fmt.Println(res)

	res2, err := c.Hello2(
		context.Background(),
		&service.HelloMessageRequest{
			Type:      "Request",
			From:      "Client go",
			Msg:       "Message from Client GO for Hello2 RPC call",
			RequestId: 101},
	)
	if err != nil {
		log.Fatalf("Can not get details line no- 36 %s", err)
	}
	fmt.Println(res2)

}
