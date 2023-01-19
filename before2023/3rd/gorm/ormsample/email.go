package ormsample

import "github.com/jinzhu/gorm"

type Email struct {
	gorm.Model
	Email  string
	UserID int
}

type EmailDao struct {
	db *gorm.DB
}

func NewEmailDao(db *gorm.DB) *EmailDao {
	return &EmailDao{
		db: db,
	}
}

func (p *EmailDao) Insert(email *Email) {
	p.db.Create(email)
}
