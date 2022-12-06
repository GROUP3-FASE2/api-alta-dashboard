package service

import (
	"api-alta-dashboard/features/class"
	"errors"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classRepository class.RepositoryInterface
	validate        *validator.Validate
}

func New(repo class.RepositoryInterface) class.ServiceInterface {
	return &classService{
		classRepository: repo,
		validate:        validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *classService) CreateClass(input class.Core) (err error) {
	// validate
	if input.ID == 0 || input.Name == "" || input.UserID == 0 {
		return errors.New("data harus diisi")
	}
	_, errCreate := service.classRepository.CreateClass(input)
	if errCreate != nil {
		return errCreate
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *classService) GetAllClass() (data []class.Core, err error) {
	data, err = service.classRepository.GetAllClass()
	return data, err
}

func (service *classService) GetByIdClass(id int) (data class.Core, err error) {
	data, err = service.classRepository.GetByIdClass(id)
	if err != nil {
		return data, err
	}
	return data, err

}

func (service *classService) UpdateClass(input class.Core, id int) error {
	// if input.Password != "" {
	// 	generate, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	// 	input.Password = string(generate)
	// }

	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.classRepository.GetByIdClass(id)
	if errFindId != nil {
		return errFindId
	}

	// proses
	_, err := service.classRepository.UpdateClass(input, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *classService) DeleteClass(id int) error {
	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.classRepository.GetByIdClass(id)
	if errFindId != nil {
		return errFindId
	}

	// proses
	_, err := service.classRepository.DeleteClass(id)
	if err != nil {
		return err
	}
	return nil
}
