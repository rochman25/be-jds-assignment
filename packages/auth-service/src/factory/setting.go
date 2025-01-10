package factory

import (
	"auth-service/database"
	"auth-service/src/repository"
)

func (f *Factory) setMysql() {
	db := database.GetConnection()

	f.db = db
}

func (f *Factory) setUserRepository() {
	f.UserRepository = repository.NewUserRepository(f.db)
}
