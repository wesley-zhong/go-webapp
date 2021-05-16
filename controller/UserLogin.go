package controller

import (
	"fmt"

	"netease.com/reqs"
)

type UserLoginService struct {
}

func (*UserLoginService) UserLogin(login *reqs.Login) interface{} {
	//log.Info("abdd")
	fmt.Println("++++++++++++++++++++++++++ call userLogin  userName = " + login.User)
	login.User = "hello:" + login.User
	return login
}
