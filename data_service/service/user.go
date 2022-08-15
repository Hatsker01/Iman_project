package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	pb "github.com/Hatsker01/Iman_project/data_service/genproto"


	l "github.com/Hatsker01/Iman_project/data_service/pkg/logger"


	cl "github.com/Hatsker01/Iman_project/data_service/service/grpc_client"
	"github.com/Hatsker01/Iman_project/data_service/storage"
	"github.com/jmoiron/sqlx"
)

type DataService struct {
	storage storage.IStorage
	logger  l.Logger
	client  cl.GrpcClientI
}

func NewDataService(db *sqlx.DB, log l.Logger,client cl.GrpcClientI) *DataService {
	return &DataService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

var client *http.Client

type All struct {
	Meta struct {
		Pagination struct {
			Total int `json:"total"`
			Pages int `json:"pages"`
			Page  int `json:"page"`
			Limit int `json:"limit"`
			Links struct {
				Previous interface{} `json:"previous"`
				Current  string      `json:"current"`
				Next     string      `json:"next"`
			} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
	Data []struct {
		ID     int    `json:"id"`
		UserID int    `json:"user_id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	} `json:"data"`
}

func (s *DataService) CreatePost(ctx context.Context, req *pb.Link) (*pb.EmptyResp, error) {
	client = &http.Client{Timeout: 10 * time.Second}
	posts, _ := GetPosts(req.Url)

	for _, j := range posts.Posts {

		_, err := s.storage.User().CreatePost(j)

		if err != nil {
			s.logger.Error("Error get post", l.Error(err))
			return nil, err
		}
	}
	return &pb.EmptyResp{}, nil
}

func GetPosts(url string) (*pb.Pages, error) {

	var posts All
	var all pb.Pages

	err := GetJson(url, &posts)
	if err != nil {
		fmt.Printf("error getting posts: %s\n", err.Error())
	}

	for posts.Meta.Pagination.Page <= 50 {

		err := GetJson(url, &posts)

		if err != nil {
			fmt.Printf("error getting posts: %s\n", err.Error())
		} else {

			for j := 0; j <= 9; j++ {
				var post pb.Post

				post.Id = strconv.Itoa(posts.Data[j].ID)
				post.UserId = strconv.Itoa(posts.Data[j].UserID)
				post.Title = posts.Data[j].Title
				post.Body = posts.Data[j].Body

				all.Posts = append(all.Posts, &post)
				if posts.Meta.Pagination.Page >= 50 {
					break
				}
			}

		}
		url = posts.Meta.Pagination.Links.Next

	}
	return &all, nil
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
