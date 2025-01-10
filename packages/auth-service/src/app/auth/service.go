package auth

import (
	"auth-service/src/dto"
	"auth-service/src/factory"
	"auth-service/src/model"
	"auth-service/src/repository"
	"auth-service/util"
	"context"
	"strings"
)

type service struct {
	UserRepository repository.UserInterface
}

type Service interface {
	RegisterService(ctx context.Context, payload dto.RegisterRequest) (res *dto.RegisterResponse, err error)
	LoginService(ctx context.Context, payload dto.LoginRequest) (res *dto.LoginResponse, err error)
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
	plainPassword, err := util.GeneratePassword(6)
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

func (s service) LoginService(ctx context.Context, payload dto.LoginRequest) (res *dto.LoginResponse, err error) {
	//get user by nik
	user, err := s.UserRepository.FindOne(ctx, "id,nik,role,password", "nik = ?", payload.Nik)
	if err != nil {
		return nil, err
	}
	//check if nik exist
	if user.Id == 0 {
		return nil, util.NIKNOTFOUND
	}
	//validate password
	checkPass := util.CheckHashPassword(payload.Password, user.Password)
	if !checkPass {
		return nil, util.PASSWORDWRONG
	}
	//generate access token
	accessToken, err := util.CreateAccessToken(user.Id)
	if err != nil {
		return nil, err
	}

	response := dto.LoginResponse{
		ID:          user.Id,
		Nik:         user.Nik,
		Role:        user.Role,
		AccessToken: *accessToken,
	}

	res = &response
	return res, nil
}
