package ormsample

import (
	"github.com/figoxu/Figo"
	"testing"
)

func TestNewUserDao_Insert(t *testing.T) {
	userDao := NewUserDao(env.db)
	userDao.Insert(&User{
		Name: "figo",
	})
}

func TestUserDao_GetById(t *testing.T) {
	userDao := NewUserDao(env.db)
	user := userDao.GetById(1)
	Figo.PrintJson("-->", user)
}
