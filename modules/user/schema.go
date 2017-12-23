package user

import "github.com/jinzhu/gorm"

type RegisterRequest struct {
	FirstName string `json:"firstname" xml:"firstname" form:"firstname" query:"firstname" valid:"required"`
	LastName  string `json:"lastname" xml:"lastname" form:"lastname" query:"lastname" valid:"required"`
	Email     string `json:"email" xml:"email" form:"email" query:"email" valid:"email,required"`
	Password  string `json:"password" xml:"password" form:"password" query:"password"  valid:"required,length(6|50)"`
}

type LoginRequest struct {
	Email    string `json:"email" xml:"email" form:"email" query:"email" valid:"email,required"`
	Password string `json:"password" xml:"password" form:"password" query:"password"  valid:"required,length(6|50)"`
}

type UserSchema struct {
	gorm.Model
	FirstName      string
	LastName       string
	Email          string `gorm:"not null;unique"`
	Password       string `gorm:"not null"`
	Roles          string
	IsActivated    bool
	ActivationCode string
	ResetCode      string
}

func (UserSchema) TableName() string {
	return "users"
}
