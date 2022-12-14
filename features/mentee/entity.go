package mentee

import (
	_class "api-alta-dashboard/features/class"
	"time"
)

type Core struct {
	ID                uint
	Name              string
	Status            string
	Gender            string
	Nickname          string
	Address           string `validate:"required"`
	HomeAddress       string
	Email             string `validate:"required,email"`
	Telegram          string `validate:"required"`
	Discord           string
	Phone             string `validate:"required"`
	EmergencyName     string `validate:"required"`
	EmergencyPhone    string `validate:"required"`
	EmergencyStatus   string `validate:"required"`
	EducationType     string `validate:"required"`
	EducationMajor    string
	EducationGraduate string
	ClassID           uint `validate:"required"`
	Logs              []LogCore
	Class             _class.Core
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type LogCore struct {
	ID        uint
	Title     string
	Feedback  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      UserCore
	Mentee    Core
}

type UserCore struct {
	ID       uint
	FullName string
	Role     string
	Status   string
}

type ServiceInterface interface {
	GetAll(queryName, queryStatus, queryIdClass, queryEdType string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(queryName, queryStatus, queryIdClass, queryEdType string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	FindUser(email string) (data Core, err error)
}
