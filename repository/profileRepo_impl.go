package repository

type ProfileRepositoryImpl struct {
}

func (p *ProfileRepositoryImpl) Hello(name string) string {
	return "hello" + name
}
