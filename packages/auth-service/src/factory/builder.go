package factory

func (f *Factory) BuildRestFactory() *Factory {
	f.setMysql()

	return f
}
