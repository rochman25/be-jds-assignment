package factory

import (
	"auth-service/src/repository"
	"context"
	"gorm.io/gorm"
)

type Factory struct {
	db             *gorm.DB
	ctx            context.Context
	UserRepository repository.UserInterface
}

func NewFactory(ctx context.Context) *Factory {
	return &Factory{
		ctx: ctx,
	}
}
