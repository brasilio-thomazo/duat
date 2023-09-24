package service

import (
	"context"
	"log"

	"br.dev.optimus/duat/pb"
	"br.dev.optimus/duat/repository"
	"gorm.io/gorm"
)

type DepartmentService struct {
	pb.UnimplementedDepartmentServiceServer
	repository repository.DepartmentRepository
}

func NewDepartmentService(reader *gorm.DB, writer *gorm.DB) *DepartmentService {
	return &DepartmentService{repository: repository.NewDepartmentRepositoryDB(reader, writer)}
}

func (s *DepartmentService) FindByID(ctx context.Context, req *pb.IDRequest) (*pb.DepartmentReply, error) {
	log.Printf("receive a findById with id %d", req.Id)

	reply, err := s.repository.FindByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DepartmentService) All(ctx context.Context, in *pb.PageRequest) (*pb.DepartmentListReply, error) {
	log.Printf("receive a all with page %d", in.Page)
	reply, err := s.repository.All(ctx, in)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
