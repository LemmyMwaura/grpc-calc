package main

import (
	"log"
	"net"

	"github.com/LemmyMwaura/grpccalc/calcserver"
	"github.com/LemmyMwaura/grpccalc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("failed to create error", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterCalculatorServer(s, calcserver.NewCalcServer())

	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
