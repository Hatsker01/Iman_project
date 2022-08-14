package service

import (
	"context"

	pb "github.com/Hatsker01/Iman_project/post_serice/genproto"
	l "github.com/Hatsker01/Iman_project/post_serice/pkg/logger"
	"github.com/Hatsker01/Iman_project/post_serice/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
    cl "github.com/Hatsker01/Iman_project/post_serice/service/grpc_client"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
    client  cl.GrpcClientI
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger,client cl.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
        client:  client,
	}
}

func (s *UserService) GetAllPost(ctx context.Context,req *pb.Empty) (*pb.GetAllRes, error) {
	posts, err := s.storage.User().GetAllPost(req)
	if err != nil {
		s.logger.Error("failed while getting all posts", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting all posts")
	}

	return posts, nil
}

func (s *UserService) GetPost(ctx context.Context, req *pb.GetPostReq) (*pb.Posts, error) {

    post,err:=s.storage.User().GetPost(req)
    if err!=nil{
        s.logger.Error("failed while getting post by id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting post by id")
    }
	return post, nil
}

func (s *UserService) DeletePost(ctx context.Context, req *pb.DelPostReq) (*pb.Empty, error) {
    _,err:=s.storage.User().DeletePost(req)
    if err!=nil{
        s.logger.Error("failed while deleting post by id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting post by id")
    }
	return &pb.Empty{}, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.UpdatePostReq) (*pb.Posts, error) {
    post,err:=s.storage.User().Update(req)
    if err!=nil{
        s.logger.Error("failed while updating post by id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating post by id")
    }
	return post, nil
}
