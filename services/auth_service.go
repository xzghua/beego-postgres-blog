package services

import (
	"bee-go-myBlog/common"
	"bee-go-myBlog/models"
	"golang.org/x/crypto/bcrypt"
)

func UserStore(u common.RegisterRequest) (id int64,err error) {
	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return 0,err
	}

	user := models.Users{
		Name:u.Name,
		Email:u.Email,
		Password:string(hashedPassword),
	}

	return models.AddUsers(&user)
}