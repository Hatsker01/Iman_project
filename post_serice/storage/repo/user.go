package repo

import (
	pb "github.com/Hatsker01/Iman_project/post_serice/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	GetAllPost(*pb.Empty) (*pb.GetAllRes, error)
	GetPost(*pb.GetPostReq)(*pb.Posts,error)
	DeletePost(*pb.DelPostReq)(*pb.Empty,error)
	Update(*pb.UpdatePostReq)(*pb.Posts,error)
}
