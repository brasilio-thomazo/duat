package main

import (
	"context"
	"flag"
	"io"
	"log"
	"path/filepath"
	"time"

	"br.dev.optimus/duat/pb"
	"br.dev.optimus/duat/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "server address")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("connection error %v", err)
	}

	defer conn.Close()
	c := pb.NewImageServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.Upload(ctx)
	if err != nil {
		log.Fatalf("upload error %v", err)
	}

	buf := make([]byte, 1024*1024)
	f, err := utils.NewFile("g511.png", "/home/brasilio/Imagens")
	if err != nil {
		log.Fatalf("get file error %v", err)
	}

	if err := f.Open(); err != nil {
		log.Fatalf("open file error %v", err)
	}

	len := 0
	req := &pb.UploadRequest{
		ClientId:   "",
		DocumentId: "",
		ImageExt:   filepath.Ext(f.Filename)[1:],
	}

	for {
		r, err := f.File.Read(buf)
		len += r
		log.Printf("read file %d of %d", len, f.Filesize)

		if err == io.EOF {
			log.Printf("read file complete")
			break
		}

		if err != nil {
			log.Printf("read file error %v", err)
			break
		}

		req.Data = buf[:r]

		if err := stream.Send(req); err != nil {
			log.Printf("send data error %v", err)
			break
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("recv error %v", err)
	}
	log.Printf("complete: filename [%s] filesize: [%d] uploaded [%d]", res.Filename, res.Filesize, res.UploadedAt)
	cancel()
}
