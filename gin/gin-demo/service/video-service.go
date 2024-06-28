package service

import "gitlab.com/pragmaticreviews/golang-gin-poc/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

// FindAll implements VideoService.
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}

// Save implements VideoService.
func (service *videoService) Save(videos entity.Video) entity.Video {
	service.videos = append(service.videos, videos)
	return videos
}

func New() VideoService {
	return &videoService{}
}
