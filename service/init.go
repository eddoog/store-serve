package service

import "github.com/eddoog/store-serve/repository"

type Service struct{}

func InitService(
	rrepository *repository.Repository,
) *Service {
	return &Service{}
}
