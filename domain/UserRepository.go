package domain

type IUserRepository interface {
	Create(*RegisterDTO) (bool, error)
	FindById(uint64) (*User, error)
	FindByEmail(string) (*User, error)
	FindAll() (*[]User, error)
	UpdateById(uint64, *UpdateUserDTO) (bool, error)
	DeleteById(uint64) (bool, error)
	GetDataView() (*[]Instansi, error)
	GetDataFunction(*DateRangeDTO) (*[]Pemakaian, error)
	SetDataProcedure(*DateRangeDTO) (bool, error)
}