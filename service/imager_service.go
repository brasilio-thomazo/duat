package service

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"br.dev.optimus/duat/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ImageService struct {
	pb.UnimplementedImageServiceServer
}

type File struct {
	filename string
	filesize int64
	file     *os.File
}

func NewImageService() *ImageService {
	return &ImageService{}
}

func NewFile(name, path string) (*File, error) {
	if strings.IndexAny(path, "/") != 0 {
		dir, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(dir, path)
	}
	filename := filepath.Join(path, name)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return nil, err
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return &File{filename: filename, filesize: 0, file: file}, nil
}

func (f *File) Write(data []byte) (int64, error) {
	len, err := f.file.Write(data)
	if err != nil {
		return -1, err
	}
	f.filesize += int64(len)
	return int64(len), nil
}

func (s *ImageService) Store(stream pb.ImageService_StoreServer) error {
	_, err := stream.Recv()
	if err != nil {
		return status.Error(codes.Unknown, err.Error())
	}

	file, err := NewFile("test.jpg", "images")

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for {
		log.Printf("waiting to receive more data")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("no more data")
			break
		}
		data := req.Image
		_, err = file.Write(data)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		// stream.SendMsg(&pb.UploadReply{Write: file.filesize, Filename: file.filename})
		log.Printf("write %d bytes in %s", file.filesize, file.filename)
	}

	return nil
}

func (s *ImageService) Upload(stream pb.ImageService_UploadServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Error(codes.Unknown, err.Error())
	}
	log.Printf("receive upload size %d", len(req.Data))

	uid, err := uuid.NewRandom()
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	name := fmt.Sprintf("%s.%s", uid.String(), req.ImageExt)

	file, err := NewFile(name, "images")

	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	reply := &pb.UploadReply{
		Filename: file.filename,
		Filesize: 0,
	}

	data := req.Data
	r, err := file.Write(data)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	// stream.SendMsg(&pb.UploadReply{Write: file.filesize, Filename: file.filename})
	log.Printf("write %d bytes in %s", file.filesize, file.filename)
	reply.Filesize += uint64(r)

	for {
		log.Printf("waiting to receive more data")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Printf("no more data")
			reply.UploadedAt = uint64(time.Now().Unix())
			break
		}
		data := req.Data
		r, err := file.Write(data)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		reply.Filesize += uint64(r)
		// stream.SendMsg(&pb.UploadReply{Write: file.filesize, Filename: file.filename})
		log.Printf("write %d bytes in %s", file.filesize, file.filename)
	}

	err = stream.SendAndClose(reply)

	return nil
}
