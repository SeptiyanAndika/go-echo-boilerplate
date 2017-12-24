package user

import (
	"boilerplate/utils"
	"errors"
)

type (
	UserRepositoryInterface interface {
		Create(firstName, lastName, email, password string) (error, UserSchema)
		FindByEmail(email string) (error, UserSchema)
	}
	UserRepository struct {
	}
)

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

func (ur *UserRepository) Create(firstName, lastName, email, hashPassword string) (error, UserSchema) {
	user := UserSchema{FirstName: firstName, LastName: lastName, Email: email, Password: hashPassword, Roles: "user", IsActivated: true}
	err := utils.GetInstanceDB().Db.Create(&user).Error
	return err, user
}

func (ur *UserRepository) FindByEmail(email string) (error, UserSchema) {
	user := UserSchema{}
	utils.GetInstanceDB().Db.Where("email = ?", email).First(&user)
	if (UserSchema{}) == user {
		return errors.New("user not found"), UserSchema{}
	} else {
		return nil, user
	}

}
