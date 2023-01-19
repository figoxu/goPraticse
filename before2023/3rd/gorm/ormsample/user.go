package ormsample

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
	Emails     []Email
	Languages  []Language `gorm:"many2many:user_languages;"`
	Friends    []*User    `gorm:"many2many:UserFriends;association_jointable_foreignkey:friend_id"`
	Fids       IntArray   `gorm:"type:integer[]"`
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (p *UserDao) Insert(user *User) {
	p.db.Create(user)
}

func (p *UserDao) GetById(id int) *User {
	user := &User{}
	p.db.Where("id=?", id).Preload("CreditCard").Preload("Emails").Preload("Languages").Preload("Friends").Find(user)
	return user
}
