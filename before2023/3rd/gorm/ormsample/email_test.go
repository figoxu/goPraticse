package ormsample

import "testing"

func TestEmailDao_Insert(t *testing.T) {
	emailDao := NewEmailDao(env.db)
	emailDao.Insert(&Email{
		Email:  "xujh945@qq.com",
		UserID: 1,
	})
	emailDao.Insert(&Email{
		Email:  "figo@qq.com",
		UserID: 1,
	})
	emailDao.Insert(&Email{
		Email:  "xujianhui@gmail.com",
		UserID: 1,
	})
}
