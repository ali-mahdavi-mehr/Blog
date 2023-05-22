package manager

import (
	"context"
	"github.com/alima12/Blog-Go/database"
	"github.com/alima12/Blog-Go/models"
	"github.com/alima12/Blog-Go/service/compiles"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type PostService struct {
	compiles.UnimplementedPostServiceServer
}

func ConvertToTimestamp(t time.Time) (*timestamp.Timestamp, error) {
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (ps *PostService) GetSinglePost(ctx context.Context, request *compiles.RetrievePost) (*compiles.SinglePostResponse, error) {
	var post models.Post
	err := post.GetOne(request.Slug)
	if err != nil {
		errMessage := err.Error()
		return nil, status.Error(codes.NotFound, errMessage)
	}
	postCreatedTime, _ := ConvertToTimestamp(post.CreatedAt)
	postUpdatedTime, _ := ConvertToTimestamp(post.UpdatedAt)
	postStatus, _ := post.Status.Value()
	response := &compiles.SinglePostResponse{
		Title:     post.Title,
		Content:   post.Content,
		UserID:    int32(post.UserID),
		Slug:      post.Slug,
		Views:     post.Views,
		ID:        int32(post.ID),
		CreatedAt: postCreatedTime.Seconds,
		UpdatedAt: postUpdatedTime.Seconds,
		Status:    postStatus,
	}
	return response, nil
}

func (ps *PostService) GetAllPosts(context.Context, *compiles.Empty) (*compiles.AllPostResponse, error) {
	db := database.GetDB()
	var posts []models.Post
	err := db.Model(&models.Post{}).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	objects := make([]*compiles.SinglePostResponse, 0)
	for _, post := range posts {
		postCreatedTime, _ := ConvertToTimestamp(post.CreatedAt)
		postUpdatedTime, _ := ConvertToTimestamp(post.UpdatedAt)
		postStatus, _ := post.Status.Value()
		p := &compiles.SinglePostResponse{
			Title:     post.Title,
			Content:   post.Content,
			UserID:    int32(post.UserID),
			Slug:      post.Slug,
			Views:     post.Views,
			ID:        int32(post.ID),
			CreatedAt: postCreatedTime.Seconds,
			UpdatedAt: postUpdatedTime.Seconds,
			Status:    postStatus,
		}
		objects = append(objects, p)
	}
	response := &compiles.AllPostResponse{Posts: objects}
	return response, nil
}
