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

func UserPwdCheck(u common.LoginRequest) (user *models.Users, res bool) {
	user,err := models.GetUsersByName(u.Name)
	if err != nil {
		return user,false
	}

	password := []byte(u.Password)
	hashedPassword := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(hashedPassword,password)
	if err != nil {
		//密码不对
		return user,false
	}
	return user,true
}