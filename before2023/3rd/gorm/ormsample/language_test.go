package ormsample

import "testing"

func TestLanguageDao_Insert(t *testing.T) {
	languageDao := NewLanguageDao(env.db)
	languageDao.Insert(&Language{
		Name:"CN",
	})
	languageDao.Insert(&Language{
		Name:"EN",
	})
}
