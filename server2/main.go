package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "../proto2"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(stream pb.Greeter_SayHelloServer) error {

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		message := in.GetName()
		fmt.Println("受取：", message)

		err = stream.Send(&pb.HelloReply{Message: message + ":xxxx"})
		log.Printf("Received: %v", in.Name)

		if err != nil {
			return err
		}
		time.Sleep(time.Second * 1)

	}

}

func main() {
	// リッスン処理
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// サーバ起動
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
