package main

import (
	"github.com/figoxu/Figo"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestUserBo(t *testing.T) {
	userDao := UserDao{}
	user := User{}
	userDao.GetById(sqlite_db, 1, &user)
	bo := UserBo{
		Base:  user,
		BoVal: "这是Bo测试",
	}
	logrus.WithField("bo", Figo.JsonString(bo)).Println(">>")
}
