package ormsample

import (
	"github.com/figoxu/Figo"
	"log"
	"testing"
)

func TestUserLanguageDao_Insert(t *testing.T) {
	userLanguageDao, userDao, languageDao := NewUserLanguageDao(env.db), NewUserDao(env.db), NewLanguageDao(env.db)
	printInfo := func() {
		user := userDao.GetById(1)
		log.Println(Figo.JsonString(user))
		language := languageDao.GetById(1)
		log.Println(Figo.JsonString(language))
	}
	printInfo()
	userLanguageDao.Insert(&UserLanguage{
		UserId:1,
		LanguageId:1,
	})
	userLanguageDao.Insert(&UserLanguage{
		UserId:1,
		LanguageId:2,
	})
	userLanguageDao.Insert(&UserLanguage{
		UserId:2,
		LanguageId:1,
	})
	userLanguageDao.Insert(&UserLanguage{
		UserId:2,
		LanguageId:2,
	})
	log.Println("------>")
	printInfo()

}

func TestUserLanguageDao_Delete(t *testing.T) {userLanguageDao, userDao, languageDao := NewUserLanguageDao(env.db), NewUserDao(env.db), NewLanguageDao(env.db)
	printInfo := func() {
		user := userDao.GetById(1)
		log.Println(Figo.JsonString(user))
		language := languageDao.GetById(1)
		log.Println(Figo.JsonString(language))
	}
	printInfo()
	userLanguageDao.Delete(&UserLanguage{
		UserId:1,
		LanguageId:1,
	})
	userLanguageDao.Delete(&UserLanguage{
		UserId:1,
		LanguageId:2,
	})
	userLanguageDao.Delete(&UserLanguage{
		UserId:2,
		LanguageId:1,
	})
	userLanguageDao.Delete(&UserLanguage{
		UserId:2,
		LanguageId:2,
	})
	log.Println("------>")
	printInfo()
}
