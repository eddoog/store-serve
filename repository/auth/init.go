package auth

type IAuthRepository interface{}

type AuthRepository struct{}

func InitAuthRepository() IAuthRepository {
	return &AuthRepository{}
}
