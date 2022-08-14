package service

import (
	"fmt"

	"github.com/Hatsker01/Iman_project/api_getaway/config"
	pb "github.com/Hatsker01/Iman_project/api_getaway/genproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	PostService() pb.PostServiceClient
	DataService() pb.DataServiceClient
}

type serviceManager struct {
	postService pb.PostServiceClient
	dataService pb.DataServiceClient
}

func (s *serviceManager) PostService() pb.PostServiceClient {
	return s.postService
}
func (s *serviceManager) DataService() pb.DataServiceClient {
	return s.dataService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostSericeHost, conf.PostSericePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	connData, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.DataServiceHost, conf.DataServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		postService: pb.NewPostServiceClient(connPost),
		dataService: pb.NewDataServiceClient(connData),
	}

	return serviceManager, nil
}
