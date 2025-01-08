package factory

import "auth-service/database"

func (f *Factory) setMysql() {
	db := database.GetConnection()

	f.db = db
}
