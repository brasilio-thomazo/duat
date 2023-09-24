package repository

import (
	"context"

	"br.dev.optimus/duat/model"
	"br.dev.optimus/duat/pb"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type DepartmentRepository interface {
	FindByID(ctx context.Context, in *pb.IDRequest) (*pb.DepartmentReply, error)
	All(ctx context.Context, in *pb.PageRequest) (*pb.DepartmentListReply, error)
}

type DepartmentRepositoryDB struct {
	DepartmentRepository
	writer *gorm.DB
	reader *gorm.DB
}

func NewDepartmentRepositoryDB(reader *gorm.DB, writer *gorm.DB) *DepartmentRepositoryDB {
	return &DepartmentRepositoryDB{writer: writer, reader: reader}
}

func (r *DepartmentRepositoryDB) FindByID(ctx context.Context, in *pb.IDRequest) (*pb.DepartmentReply, error) {
	department := &model.Department{}
	reply := &pb.DepartmentReply{}

	result := r.reader.WithContext(ctx).Where("id = ?", in.Id).First(&department)
	if result.Error != nil {
		return nil, status.Error(codes.NotFound, result.Error.Error())
	}

	err := copier.Copy(&reply, &department)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return reply, nil
}

func (r *DepartmentRepositoryDB) All(ctx context.Context, in *pb.PageRequest) (*pb.DepartmentListReply, error) {
	var departments []*model.Department
	var list []*pb.DepartmentReply

	result := r.reader.WithContext(ctx).Find(&departments)

	if result.Error != nil {
		return nil, status.Error(codes.Internal, result.Error.Error())
	}

	err := copier.Copy(&list, &departments)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	reply := &pb.DepartmentListReply{List: list}
	return reply, nil
}
