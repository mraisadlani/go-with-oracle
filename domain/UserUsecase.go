package domain

type IUserUsecase interface {
	CreateUser(*RegisterDTO) (bool, error)
	FindUser(uint64) (*User, error)
	ViewAll() (*[]User, error)
	UpdateUser(uint64, *UpdateUserDTO) (bool, error)
	DeleteUser(uint64) (bool, error)
	GetDataByView() (*[]Instansi, error)
	GetDataByFunction(*DateRangeDTO) (*[]Pemakaian, error)
	SetDataByProcedure(*DateRangeDTO) (bool, error)
}