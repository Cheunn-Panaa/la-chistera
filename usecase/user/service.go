package user

import (
	"github.com/cheunn-panaa/la-chistera/entity"

	log "github.com/sirupsen/logrus"
)

type Service struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) *Service {
	log.Info("Initialising User Service...")
	return &Service{
		repository: r,
	}
}

func (s *Service) CreateUser(user *entity.User) (*entity.User, error) {
	e, err := entity.NewUser(user.Email, user.Password, user.FirstName, user.LastName)
	if err != nil {
		return e, err
	}
	return s.repository.Create(user)
}
func (s *Service) ListUsers() (*[]entity.User, error) {
	return nil, nil
}
