package postgres

import (
	pb "github.com/Hatsker01/Iman_project/data_service/genproto"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreatePost(post *pb.Post) (*pb.Post, error) {

	var retpost pb.Post

	queryInsertPost := `INSERT INTO posts(id,user_id,title,body) VALUES ($1,$2,$3,$4) RETURNING id,user_id,title,body`
	err := r.db.QueryRow(queryInsertPost, post.Id, post.UserId, post.Title, post.Body).Scan(
		&retpost.Id,
		&retpost.UserId,
		&retpost.Title,
		&retpost.Body,
	)
	if err != nil {
		return nil, err
	}

	return &retpost, nil
}
