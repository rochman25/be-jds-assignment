package repository

import (
	"auth-service/src/model"
	"context"
	"gorm.io/gorm"
)

type UserInterface interface {
	Finds(ctx context.Context, selectField string, query string, args ...interface{}) (response []model.User, err error)
	FindOne(ctx context.Context, selectField string, query string, args ...interface{}) (response *model.User, err error)
	Store(ctx context.Context, data *model.User) error
}

type user struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (r *user) Finds(ctx context.Context, selectField string, query string, args ...interface{}) (response []model.User, err error) {
	err = r.Db.WithContext(ctx).
		Model(&model.User{}).
		Select(selectField).
		Where(query, args).
		Find(&response).
		Error
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *user) FindOne(ctx context.Context, selectField string, query string, args ...interface{}) (response *model.User, err error) {
	err = r.Db.WithContext(ctx).
		Model(&model.User{}).
		Select(selectField).
		Where(query, args...).
		Find(&response).
		Error
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (r *user) Store(ctx context.Context, data *model.User) error {
	tx := r.Db.WithContext(ctx)
	if err := tx.Model(model.User{}).Create(&data).Error; err != nil {
		return err
	}
	return nil
}
