package service_models

import (
	"context"
	"errors"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/service/compiles"
)

type PostService struct {
	compiles.UnimplementedPostServiceServer
}

func (ps *PostService) GetSinglePost(context.Context, *compiles.RetrievePost) (*compiles.SinglePostResponse, error) {
	db := database.GetDB()
	var post models.Post
	result := db.First(&post, 1)
	if result.Error != nil {
		return nil, errors.New("post not found")
	}
	response := &compiles.SinglePostResponse{
		Title:  post.Title,
		Body:   post.Body,
		UserID: int32(post.UserID),
	}
	return response, nil
}
