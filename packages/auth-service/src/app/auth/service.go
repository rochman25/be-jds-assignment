package auth

import (
	"auth-service/src/dto"
	"auth-service/src/factory"
	"auth-service/src/model"
	"auth-service/src/repository"
	"auth-service/util"
	"context"
	"crypto/rand"
	"math/big"
	"strings"
)

type service struct {
	UserRepository repository.UserInterface
}

type Service interface {
	RegisterService(ctx context.Context, payload dto.RegisterRequest) (res *dto.RegisterResponse, err error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s service) RegisterService(ctx context.Context, payload dto.RegisterRequest) (res *dto.RegisterResponse, err error) {
	//get user by nik
	user, err := s.UserRepository.FindOne(ctx, "id", "nik = ?", payload.Nik)
	if err != nil {
		return nil, err
	}
	//check if nik exist
	if user.Id > 0 {
		return nil, util.NIKALREADYEXIST
	}
	//generate password
	plainPassword, err := generatePassword(6)
	if err != nil {
		return nil, err
	}

	//generate hash password
	password := util.GenerateHashPassword(plainPassword)

	userData := model.User{
		Nik:      payload.Nik,
		Role:     strings.ToLower(payload.Role),
		Password: password,
	}
	err = s.UserRepository.Store(ctx, &userData)
	if err != nil {
		return nil, err
	}

	response := dto.RegisterResponse{
		Nik:      payload.Nik,
		Role:     payload.Role,
		Password: plainPassword,
	}

	res = &response

	return res, nil
}

// generatePassword generates a random alphanumeric password
func generatePassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}
	return string(password), nil
}
