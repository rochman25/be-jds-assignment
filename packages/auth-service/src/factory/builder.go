package factory

func (f *Factory) BuildRestFactory() *Factory {
	f.setMysql()
	f.setUserRepository()

	return f
}
