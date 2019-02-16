package ormsample

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
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
	p.db.Where("id=?", id).Preload("CreditCard").Find(user)
	return user
}
