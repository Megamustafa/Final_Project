package services

import (
	"aquaculture/models"
	"aquaculture/repositories"
	"aquaculture/utils"
)

type UserService struct {
	Repository repositories.UserRepository
	jwtOptions models.JWTOptions
}

func InitUserService(jwtOptions models.JWTOptions) UserService {
	return UserService{
		Repository: &repositories.UserRepositoryImpl{},
		jwtOptions: jwtOptions,
	}
}

func (us *UserService) Register(registerReq models.RegisterRequest) (models.User, error) {
	return us.Repository.Register(registerReq)
}

func (us *UserService) LoginAdmin(loginReq models.LoginRequest) (string, error) {
	admin, err := us.Repository.GetByEmailAdmin(loginReq)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(int(admin.ID), us.jwtOptions)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) LoginUser(loginReq models.LoginRequest) (string, error) {
	user, err := us.Repository.GetByEmailUser(loginReq)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(int(user.ID), us.jwtOptions)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (us *UserService) GetUserInfo(id string) (models.User, error) {
	return us.Repository.GetUserInfo(id)
}

func (us *UserService) GetAdminInfo(id string) (models.Admin, error) {
	return us.Repository.GetAdminInfo(id)
}
