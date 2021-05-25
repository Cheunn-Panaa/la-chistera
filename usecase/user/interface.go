package user

import (
	"github.com/cheunn-panaa/la-chistera/entity"
)

//Reader interface
type Reader interface {
	Get(id uint) (*entity.User, error)
	List() (*[]entity.User, error)
}

//Writer user writer
type Writer interface {
	Create(user *entity.User) (*entity.User, error)
	Update(e *entity.User) error
	Delete(id uint) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	ListUsers() (*[]entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	/* UpdateUser(user *entity.User) error
	GetUser(id uint) (*entity.User, error)
	DeleteUser(id uint) error */
}
