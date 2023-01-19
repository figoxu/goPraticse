package main

import (
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestUserDao(t *testing.T) {
	userDao := UserDao{}
	userDao.Insert(sqlite_db, &User{
		Name:  "figo",
		Age:   18,
		Phone: "18888888888",
	})
	userDao.Insert(sqlite_db, &User{
		Name:  "andy",
		Age:   18,
		Phone: "18666666666",
	})
	user := User{}
	userDao.GetById(sqlite_db, 1, &user)
	logrus.WithField("user", Figo.JsonString(user)).Println("print")
	user.Name = "lucy"
	user.Phone = "18500589946"
	utee.Chk(userDao.Update(sqlite_db, &user, "name", "phone"))
	userDao.GetById(sqlite_db, 1, &user)
	logrus.WithField("user", Figo.JsonString(user)).Println("print")
}
