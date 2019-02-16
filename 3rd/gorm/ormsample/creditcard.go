package ormsample

import "github.com/jinzhu/gorm"

type CreditCard struct {
	gorm.Model
	UserID int
	Number string
}

type CreditCardDao struct {
	db *gorm.DB
}

func NewCreditCardDao(db *gorm.DB) *CreditCardDao {
	return &CreditCardDao{
		db: db,
	}
}

func(p *CreditCardDao) Insert(creditCard *CreditCard){
	p.db.Create(creditCard)
}

