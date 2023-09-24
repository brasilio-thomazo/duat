package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"br.dev.optimus/duat/database"
	"br.dev.optimus/duat/pb"
	"br.dev.optimus/duat/service"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "server port")

func main() {
	flag.Parse()

	dsn := "host=localhost user=postgres password=postgres dbname=client_dev port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	reader, err := database.NewReader(dsn)
	if err != nil {
		log.Fatalf("connect to database reader server error %v", err)
	}

	writer, err := database.NewWriter(dsn)
	if err != nil {
		log.Fatalf("connect to database writer server error %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("listen %d error [%v]", *port, err)
	}

	s := grpc.NewServer()
	sDepartment := service.NewDepartmentService(reader, writer)
	sImage := service.NewImageService()

	pb.RegisterDepartmentServiceServer(s, sDepartment)
	pb.RegisterImageServiceServer(s, sImage)

	log.Printf("server listening %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server start falied %v", err)
	}
}
