package repository

import (
	"github.com/cheunn-panaa/la-chistera/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepo will create an implementation of user.Repository
func NewUserRepo(db *gorm.DB) *UserRepository {

	log.Info("Initialising User Repository...")
	return &UserRepository{
		DB: db,
	}
}

//Create an user
func (r *UserRepository) Create(user *entity.User) (*entity.User, error) {
	result := r.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

//Get an user
func (r *UserRepository) Get(id uint) (*entity.User, error) {
	var user entity.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

//Update an user
func (r *UserRepository) Update(e *entity.User) error {

	return nil
}

//List users
func (r *UserRepository) List() (*[]entity.User, error) {
	var userList []entity.User
	result := r.DB.Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userList, nil
}

//Delete an user
func (r *UserRepository) Delete(id uint) error {
	result := r.DB.Delete(&entity.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
