package main

import (
	"context"

	pb "./testalert"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SendAlert(ctx context.Context, in *pb.SendAlertInput) (*pb.Empty, error) {
	if in.Code == pb.SendAlertInput_Promote_Success {
		return nil, nil
	}
	return nil, nil
}

func main() {
	return nil
}
