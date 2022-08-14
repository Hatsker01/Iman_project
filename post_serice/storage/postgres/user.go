package postgres

import (
	pb "github.com/Hatsker01/Iman_project/post_serice/genproto"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) GetAllPost(not *pb.Empty) (*pb.GetAllRes, error) {
	var posts []*pb.Posts

	queryPostsGet := `SELECT id,user_id,title,body from posts`
	rows, err := r.db.Query(queryPostsGet)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post pb.Posts
		err := rows.Scan(
			&post.Id,
			&post.UserId,
			&post.Title,
			&post.Body,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return &pb.GetAllRes{
		Posts: posts,
	}, nil
}

func (r *userRepo) GetPost(id *pb.GetPostReq) (*pb.Posts, error) {
	var post pb.Posts

	queryGetbyId := `SELECT id,user_id,title,body from posts where id=$1`
	err := r.db.QueryRow(queryGetbyId, id.Id).Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Body,
	)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *userRepo) DeletePost(id *pb.DelPostReq) (*pb.Empty, error) {
	queryDel := `DELETE from posts where id=$1`
	_, err := r.db.Exec(queryDel, id.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (r *userRepo) Update(post *pb.UpdatePostReq) (*pb.Posts, error) {

	var newpost pb.Posts

	queryUpdate := `UPDATE posts SET title=$2,body=$3 where id = $1 RETURNING id,user_id,title,body`
	err := r.db.QueryRow(queryUpdate, post.Id, post.Title, post.Body).Scan(
		&newpost.Id,
		&newpost.UserId,
		&newpost.Title,
		&newpost.Body,
	)
	if err != nil {
		return nil, err
	}

	return &newpost, nil
}
