package main

import (
	"context"
	"log"
	"net"
	"testing"

	"go.uber.org/goleak"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	pb "grpc-testing-with-bufconn/proto/gen/proto"
)

const bufSize = 1024 * 1024

func TestSayHello(t *testing.T) {
	t.Cleanup(func() {
		goleak.VerifyNone(t)
	})
	lis := bufconn.Listen(bufSize)
	t.Cleanup(func() {
		lis.Close()
	})
	s := grpc.NewServer()
	t.Cleanup(func() {
		s.Stop()
	})
	pb.RegisterGreeterServer(s, &server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	ctx := t.Context()
	conn, err := grpc.NewClient("passthrough://bufnet",
		grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("Failed to create new client: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "test"})
	if err != nil {
		t.Fatal(err)
	}

	if resp.GetMessage() != "Hello test" {
		t.Fatal("hello reply must be 'Hello test'")
	}
}
