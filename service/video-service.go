package service

import (
	"go-api/entity"
)

type VideoService interface {
	SaveVideo(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func NewVideoService() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

func (s *videoService) SaveVideo(video entity.Video) entity.Video {
	s.videos = append(s.videos, video)

	return video
}

func (s *videoService) FindAll() []entity.Video {
	return s.videos
}
