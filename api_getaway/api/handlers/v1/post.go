package v1

import (
	"context"

	//"encoding/json"
	"fmt"
	"net/http"
	"time"

	pb "github.com/Hatsker01/Iman_project/api_getaway/genproto"
	l "github.com/Hatsker01/Iman_project/api_getaway/pkg/logger"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

//GetPost by id
//@Summary Get post summary
//Description This api for getting post by id
//@Tags post
//@Accept json
//@Produce json
//@Param id path string true "Post_id"
//@Success 200 {string} Post!
//@Router /v1/users/{id} [get]
func (h *handlerV1) GetPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetPost(
		ctx, &pb.GetPostReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return

	}
	c.JSON(http.StatusOK, response)

}

// CreatePost godoc
// @Summary Create new post
// @Description This API for creating a new post
// @Tags Post
// @Accept json
// @Param body body Link true "body"
// @Produce json
// @Success 200
// @Router /v1/posts [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		body     pb.Link   
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.DataService().CreatePost(ctx, &body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create post", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

//Delete Post by id
//@Summary Delete post summary
//Description This api for delete post by id
//@Tags user
//@Accept json
//@Produce json
//@Param id path string true "Post_id"
//@Success 200 Succesfully deleted Post!
//@Router /v1/users/delete/{id} [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeletePost(
		ctx, &pb.DelPostReq{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return

	}
	c.JSON(http.StatusOK, response)

}

//UpdateUser updates post
//@Summary update post summary
//Description This api for updating post
//@Tags user
//@Accept json
//@Produce json
//@Param post body Post true "post body"
//@Success 200 Succesfully {string} Post!
//@Router /v1/users/update [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		body        pb.UpdatePostReq
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to blind json", l.Error(err))
		return
	}
	fmt.Println(&body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().Update(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)

}

//GetAllPost gets post
//@Summary Get post summary
//Description This api for getting POSTS
//@Tags post
//@Accept json
//@Produce json
//@Success 200 {string} []Post!
//@Router /v1/users/allpost [get]
func (h *handlerV1) GetAllPost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetAllPost(
		ctx, &pb.Empty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return

	}
	c.JSON(http.StatusOK, response)

}

type Post struct {
	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`
	Body   string `protobuf:"bytes,4,opt,name=body,proto3" json:"body"`
}
type CreatedPost struct {
	Link string `json:"link"`
}
type Link struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url"`
	
}
