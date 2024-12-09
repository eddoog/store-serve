package repository

import "github.com/eddoog/store-serve/repository/auth"

type Repository struct {
	AuthRepository auth.IAuthRepository
}

func InitRepository() *Repository {
	return &Repository{
		AuthRepository: auth.InitAuthRepository(),
	}
}
