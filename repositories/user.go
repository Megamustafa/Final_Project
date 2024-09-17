package repositories

import (
	"aquaculture/database"
	"aquaculture/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryImpl struct{}

func InitUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) Register(registerReq models.RegisterRequest) (models.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err
	}

	var user models.User = models.User{
		Email:       registerReq.Email,
		Password:    string(password),
		FullName:    registerReq.FullName,
		Address:     registerReq.Address,
		PhoneNumber: registerReq.PhoneNumber,
	}

	result := database.DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	err = result.Last(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetByEmailUser(loginReq models.LoginRequest) (models.User, error) {
	var user models.User

	err := database.DB.First(&user, "email = ?", loginReq.Email).Error

	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetByEmailAdmin(loginReq models.LoginRequest) (models.Admin, error) {
	var admin models.Admin

	err := database.DB.First(&admin, "email = ?", loginReq.Email).Error

	if err != nil {
		return models.Admin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginReq.Password))

	if err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}

func (ur *UserRepositoryImpl) GetUserInfo(id string) (models.User, error) {
	var user models.User

	err := database.DB.First(&user, "id = ?", id).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) GetAdminInfo(id string) (models.Admin, error) {
	var admin models.Admin

	err := database.DB.First(&admin, "id = ?", id).Error

	if err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}
