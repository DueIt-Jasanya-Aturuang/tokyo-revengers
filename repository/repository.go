package repository

type AllServiceRepoImpl struct {
	Profile *ProfileRepositoryImpl
}

func NewAllServiceRepoImpl() *AllServiceRepoImpl {
	return &AllServiceRepoImpl{
		Profile: &ProfileRepositoryImpl{},
	}
}
