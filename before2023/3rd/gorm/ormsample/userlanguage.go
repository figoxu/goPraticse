package ormsample

import "github.com/jinzhu/gorm"

type UserLanguage struct {
	UserId     uint
	LanguageId uint
}

type UserLanguageDao struct {
	db *gorm.DB
}

func NewUserLanguageDao(db *gorm.DB) *UserLanguageDao {
	return &UserLanguageDao{
		db: db,
	}
}

func (p *UserLanguageDao) Insert(ul *UserLanguage) {
	p.db.Create(ul)
}

func (p *UserLanguageDao) Delete(ul *UserLanguage) {
	p.db.Delete(ul)
}
