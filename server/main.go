package main

import (
	"context"
	"flag"
	"github.com/Manav-Aggarwal/herd-mentality/server/pb/hm"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	questionsPtr := flag.String("questions", "", "Path to file with questions")
	laddrPtr := flag.String("laddr", "0.0.0.0:9988", "Listen address (tcp/ip)")

	flag.Parse()

	log.Println("questions file:", *questionsPtr)
	log.Println("laddr:", *laddrPtr)

	lis, err := net.Listen("tcp", *laddrPtr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	srv := &server{}
	srv.endpoint = "172.16.24.190:36657"
	file, err := os.Open(*questionsPtr)
	if err != nil {
		log.Fatalf("cannot load questions file: %v", err)
	}
	srv.loadQuestions(file)
	log.Printf("# of loaded questions: %d", len(srv.questions))

	srv.start(context.Background())
	hm.RegisterHerdMentalityServer(grpcServer, srv)

	grpcServer.Serve(lis)
}
