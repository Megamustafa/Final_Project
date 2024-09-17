package repositories_test

import (
	"aquaculture/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		userRepository.On("Register", models.RegisterRequest{}).Return(models.User{}, nil).Once()

		result, err := userService.Register(models.RegisterRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Register | Invalid", func(t *testing.T) {
		userRepository.On("Register", models.RegisterRequest{}).Return(models.User{}, errors.New("something went wrong")).Once()

		result, err := userService.Register(models.RegisterRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByEmailUser(t *testing.T) {
	t.Run("GetByEmailUser | Valid", func(t *testing.T) {
		userRepository.On("GetByEmailUser", models.LoginRequest{}).Return(models.User{}, nil).Once()

		result, err := userService.LoginUser(models.LoginRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByEmailUser | Invalid", func(t *testing.T) {
		userRepository.On("GetByEmailUser", models.LoginRequest{}).Return(models.User{}, errors.New("something went wrong")).Once()

		result, err := userService.LoginUser(models.LoginRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByEmailAdmin(t *testing.T) {
	t.Run("GetByEmailAdmin | Valid", func(t *testing.T) {
		userRepository.On("GetByEmailAdmin", models.LoginRequest{}).Return(models.Admin{}, nil).Once()

		result, err := userService.LoginAdmin(models.LoginRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByEmailAdmin | Invalid", func(t *testing.T) {
		userRepository.On("GetByEmailAdmin", models.LoginRequest{}).Return(models.Admin{}, errors.New("something went wrong")).Once()

		result, err := userService.LoginAdmin(models.LoginRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetUserInfo(t *testing.T) {
	t.Run("GetUserInfo | Valid", func(t *testing.T) {
		userRepository.On("GetUserInfo", "1").Return(models.User{}, nil).Once()

		result, err := userService.GetUserInfo("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetUserInfo | Invalid", func(t *testing.T) {
		userRepository.On("GetUserInfo", "0").Return(models.User{}, errors.New("something went wrong")).Once()

		result, err := userService.GetUserInfo("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetAdminInfo(t *testing.T) {
	t.Run("GetAdminInfo | Valid", func(t *testing.T) {
		userRepository.On("GetAdminInfo", "1").Return(models.Admin{}, nil).Once()

		result, err := userService.GetAdminInfo("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAdminInfo | Invalid", func(t *testing.T) {
		userRepository.On("GetAdminInfo", "0").Return(models.Admin{}, errors.New("something went wrong")).Once()

		result, err := userService.GetAdminInfo("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}
