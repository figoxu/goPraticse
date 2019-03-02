package ormsample

import (
	"fmt"
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

func TestUserDao_InsertWithFriends(t *testing.T) {
	userDao := NewUserDao(env.db)
	user := userDao.GetById(1)
	env.db.Model(&user).Association("Friends").Append(&User{Name: "张三"}, &User{Name: "李四"})
	fmt.Println("HELLO")
}

func TestUserDao_Ids(t *testing.T) {
	ids := make([]int, 0)
	env.db.Model(&User{}).Pluck("id", &ids)
	for _, id := range ids {
		fmt.Println(id)
	}
}
