package labeler

import (
	"log"
	"net"

	labeler "github.com/dimgatz98/labeler/pkg/server"

	"google.golang.org/grpc"
)

func Serve() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	labeler.RegisterLabelerServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
