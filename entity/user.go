package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//Construct your model under entity.
type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`

	Bets []*MatchBets `gorm:"many2many:bets"`
}

// Only one struct per file should exists unless another struct is closely related with the one defined in this file.
//type DeleteRequest struct {
//	ID string `json:"id" binding:"required,gte=1"`
//}

//NewUser create a new user
func NewUser(email, password, firstName, lastName string) (*User, error) {
	u := &User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd
	/* err = u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}*/
	return u, nil
}

/* Validate validate data
func (u *User) Validate() error {
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Password == "" {
		return ErrInvalidEntity
	}

	return nil
}
*/

//ValidatePassword validate user password
func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 15)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
