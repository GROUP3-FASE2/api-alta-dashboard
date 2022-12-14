package user

import "time"

type Core struct {
	ID        uint
	FullName  string `valiidate:"required"`
	Email     string `valiidate:"required,email,unique"`
	Password  string `valiidate:"required"`
	Team      string
	Role      string `validiate:"required"`
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll(query string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	GetAllWithSearch(query string) (data []Core, err error)
	Create(input Core) error
	GetById(id int) (data Core, err error)
	Update(input Core, id int) error
	Delete(id int) error
	FindUser(email string) (data Core, err error)
}
