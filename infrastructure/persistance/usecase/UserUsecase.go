package usecase

import (
	"Training/go-crud-with-oracle/domain"
	"Training/go-crud-with-oracle/infrastructure/persistance/repository"
	"errors"
)

type CustomerUsecase struct {
	userRepo *repository.UserRepository
}

func BuildUserUsecase(userRepo *repository.UserRepository) *CustomerUsecase {
	return &CustomerUsecase{
		userRepo: userRepo,
	}
}

var _ domain.IUserUsecase = &CustomerUsecase{}

func (s *CustomerUsecase) CreateUser(registerDTO *domain.RegisterDTO) (bool, error) {
	find, _ := s.userRepo.FindByEmail(registerDTO.Email)

	if find.Email != "" {
		return false, errors.New("User already exists")
	} else {
		set, err := s.userRepo.Create(registerDTO)

		if err != nil {
			return false, err
		}

		return set, nil
	}
}

func (s *CustomerUsecase) FindUser(id uint64) (*domain.User, error) {
	find, err := s.userRepo.FindById(id)

	if err != nil {
		return nil, errors.New("User not found")
	}

	return find, nil
}

func (s *CustomerUsecase) ViewAll() (*[]domain.User, error) {
	get, err := s.userRepo.FindAll()

	if err != nil {
		return nil, errors.New("Data not found")
	}

	return get, nil
}

func (s *CustomerUsecase) UpdateUser(id uint64, updateDTO *domain.UpdateUserDTO) (bool, error) {
	find, _ := s.userRepo.FindById(id)

	if find.NamaUser == "" {
		return false, errors.New("User not found")
	} else {
		updateDTO.UpdateBy = find.NamaUser

		set, err := s.userRepo.UpdateById(uint64(find.ID), updateDTO)

		if err != nil {
			return false, err
		}

		return set, nil
	}
}

func (s *CustomerUsecase) DeleteUser(id uint64) (bool, error) {
	find, _ := s.userRepo.FindById(id)

	if find.NamaUser == "" {
		return false, errors.New("User not found")
	} else {
		get, err := s.userRepo.DeleteById(id)

		if err != nil {
			return false, err
		}

		return get, nil
	}
}

func (s *CustomerUsecase) GetDataByView() (*[]domain.Instansi, error) {
	get, err := s.userRepo.GetDataView()

	if err != nil {
		return nil, errors.New("Data not found")
	}

	return get, nil
}

func (s *CustomerUsecase) GetDataByFunction(dataRange *domain.DateRangeDTO) (*[]domain.Pemakaian, error) {
	get, err := s.userRepo.GetDataFunction(dataRange)

	if err != nil {
		return nil, errors.New("Data not found")
	}

	return get, nil
}

func (s *CustomerUsecase) SetDataByProcedure(dataRange *domain.DateRangeDTO) (bool, error) {
	get, err := s.userRepo.SetDataProcedure(dataRange)

	if err != nil {
		return false, errors.New("Data not found")
	}

	return get, nil
}