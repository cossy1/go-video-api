package entity

import (
	"time"

	"github.com/google/uuid"
)

type Video struct {
	Title       string    `json:"title" binding:"min=2,max=20,required"`
	Description string    `json:"description" binding:"required,max=200"`
	URL         string    `json:"url" binding:"required,url"`
	Author      User      `json:"author" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	UserID    uuid.UUID `json:"userId" binding:"required"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateVideoRequest struct {
	Title       string `json:"title" binding:"min=2,max=20,required"`
	Description string `json:"description" binding:"required,max=200"`
	URL         string `json:"url" binding:"required,url"`

	UserID    uuid.UUID `json:"userId"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

type UpdateVideoRequest struct {
	Title       string `json:"title" binding:"min=2,max=20,required"`
	Description string `json:"description" binding:"required,max=200"`
	URL         string `json:"url" binding:"required,url"`

	UpdatedAt time.Time `json:"updatedAt"`
}

type VideoResponse struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	URL         string         `json:"url"`
	Author      AuthorResponse `json:"author"`
	ID          uuid.UUID      `json:"id"`

	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func ToVideoResponse(video Video) *VideoResponse {
	return &VideoResponse{
		Title:       video.Title,
		Description: video.Description,
		URL:         video.URL,
		Author: AuthorResponse{
			FirstName: video.Author.FirstName,
			LastName:  video.Author.LastName,
			Age:       video.Author.Age,
			Email:     video.Author.Email,
			ID:        video.Author.ID,
			CreatedAt: video.Author.CreatedAt,
			UpdatedAt: video.Author.UpdatedAt,
		},
		ID:        video.ID,
		UpdatedAt: video.UpdatedAt,
		CreatedAt: video.CreatedAt,
	}
}
