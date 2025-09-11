package service

import (
	"errors"
	"fmt"
	"go-api/entity"
	"time"

	"gorm.io/gorm"
)

type VideoService interface {
	SaveVideo(entity.CreateVideoRequest) (entity.Video, error)
	GetAll(id string) ([]entity.Video, error)
	GetVideo(videoId string) (entity.Video, error)
	UpdateVideo(videoId string, req entity.UpdateVideoRequest) (entity.Video, error)
}

type videoService struct {
	db *gorm.DB
}

func NewVideoService(db *gorm.DB) VideoService {
	return &videoService{
		db: db,
	}
}

func (s *videoService) SaveVideo(request entity.CreateVideoRequest) (entity.Video, error) {
	video := entity.Video{
		Title:       request.Title,
		Description: request.Description,
		URL:         request.URL,
		UserID:      request.UserID,
		CreatedAt:   request.CreatedAt,
	}

	err := s.db.Create(&video).Error

	if err := s.db.Preload("Author").First(&video, "id = ?", video.ID).Error; err != nil {
		return entity.Video{}, err
	}

	return video, err
}

func (s *videoService) GetAll(id string) ([]entity.Video, error) {
	var videos []entity.Video

	err := s.db.Where("id = ?", id).Error

	if err != nil {
		return []entity.Video{}, err
	}

	if err := s.db.Find(&videos).Where("userId = ?", id).Error; err != nil {
		return []entity.Video{}, err
	}

	if err := s.db.Preload("Author").First(&videos).Error; err != nil {
		return []entity.Video{}, err
	}

	return videos, nil
}

func (s *videoService) GetVideo(videoId string) (entity.Video, error) {
	var video entity.Video

	err := s.db.Where("id = ?", videoId).Error

	if err != nil {
		return entity.Video{}, err
	}

	if err := s.db.Preload("Author").First(&video).Error; err != nil {
		return entity.Video{}, err
	}

	return video, nil
}

func (s *videoService) UpdateVideo(videoId string, req entity.UpdateVideoRequest) (entity.Video, error) {
	var video entity.Video

	err := s.db.First(&video, "id = ?", videoId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Video{}, fmt.Errorf("video with ID %s not found", videoId)
	}

	if err != nil {
		return entity.Video{}, fmt.Errorf("database error: %v", err)
	}

	update := map[string]interface{}{}

	if req.Title != "" {
		update["title"] = req.Title
	}
	if req.Description != "" {
		update["description"] = req.Description
	}

	if req.URL != "" {
		update["url"] = req.URL
	}

	if len(update) == 0 {
		return video, nil // No fields to update
	}

	update["updated_at"] = time.Now()

	err = s.db.Model(&video).Updates(update).Error
	if err != nil {
		return entity.Video{}, fmt.Errorf("failed to update video: %v", err)
	}

	if err := s.db.Preload("Author").First(&video).Error; err != nil {
		return entity.Video{}, err
	}

	return video, nil
}
