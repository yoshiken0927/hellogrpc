package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

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

	stdin := bufio.NewScanner(os.Stdin)

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("エラー: %v", err)
		}
		log.Printf("サーバから：%s", in.Message)

		// お返し
		stdin.Scan()
		cmd := stdin.Text()
		stream.Send(&pb.HelloRequest{Name: cmd})
	}

	stream.CloseSend()
}
