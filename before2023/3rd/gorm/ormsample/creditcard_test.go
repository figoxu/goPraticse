package ormsample

import "testing"

func TestNewCreditCardDao(t *testing.T) {
	creditCardDao := NewCreditCardDao(env.db)
	creditCardDao.Insert(&CreditCard{
		UserID:1,
		Number:"367808198411222029",
	})
}
