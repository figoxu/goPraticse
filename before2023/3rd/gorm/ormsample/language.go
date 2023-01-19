package ormsample

import "github.com/jinzhu/gorm"

type Language struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_languages;"`
}

type LanguageDao struct {
	db *gorm.DB
}

func NewLanguageDao(db *gorm.DB) *LanguageDao {
	return &LanguageDao{
		db: db,
	}
}

func (p *LanguageDao) Insert(language *Language) {
	p.db.Create(language)
}

func (p *LanguageDao) GetById(id int) (*Language) {
	language := &Language{}
	p.db.Preload("Users").Find(language, p.db.Where("id=?", id))
	return language
}
