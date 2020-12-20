package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "../proto2"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	stream, err := client.SayHello(context.Background())
	if err != nil {
		return
	}

	err = stream.Send(&pb.HelloRequest{Name: "こんにちは"})
	if err != nil {
		return
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("エラー: %v", err)
			}
			log.Printf("サーバから：%s", in.Message)

			// お返し
			stream.Send(&pb.HelloRequest{
				Name: time.Now().Format("2006-01-02 15:04:05"),
			})
		}
	}()
	<-waitc

	if err != nil {
		return
	}
	stream.CloseSend()
}
