package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"log"

	graphModel "github.com/alima12/Blog-Go/cmd/graph/graph/model"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input graphModel.CreatePost) (*graphModel.Post, error) {
	var post models.Post
	post.Title = input.Title
	post.Content = input.Content
	post.Slug = input.Slug
	post.ImageURL = input.ImageURL
	post.UserID = uint(1)
	post.Status = "published"
	db := database.GetDB()
	err := db.Create(&post).Error
	if err != nil {
		log.Fatal(err)
	}
	return &graphModel.Post{
		Title:    post.Title,
		Content:  post.Content,
		Slug:     post.Slug,
		UserID:   int(post.UserID),
		ImageURL: post.ImageURL,
		Views:    int(post.Views),
	}, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*graphModel.Post, error) {
	db := database.GetDB()
	var posts []models.Post
	err := db.Model(&models.Post{}).Order("views desc").Find(&posts).Error
	if err != nil {
		log.Fatal(err)
	}
	response := make([]*graphModel.Post, 0)
	for _, post := range posts {
		response = append(response, &graphModel.Post{
			Title:    post.Title,
			Content:  post.Content,
			Slug:     post.Slug,
			UserID:   int(post.UserID),
			ImageURL: post.ImageURL,
			Views:    int(post.Views),
		})
	}
	return response, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
