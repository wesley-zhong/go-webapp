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

func (*UserLoginService) UserUpdate(userInfo *reqs.UserInfo) interface{} {
	fmt.Println("user info update username =" + userInfo.UserName)
	return userInfo
}
