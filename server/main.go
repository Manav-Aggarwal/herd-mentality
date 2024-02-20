package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Manav-Aggarwal/herd-mentality/server/pb/hm"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 12345))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	hm.RegisterHerdMentalityServer(grpcServer, &server{})

	grpcServer.Serve(lis)

}
