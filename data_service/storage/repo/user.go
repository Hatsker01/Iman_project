package repo

import (
	pb "github.com/Hatsker01/Iman_project/data_service/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	CreatePost(*pb.Post) (*pb.Post, error)
}
