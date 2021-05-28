package ut

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Chk(err error) {
	if err != nil {
		logrus.WithField("error", fmt.Sprintf("%+v", err)).Panicln("panic ", err.Error())
	}
}

func HasErrFeature(err error, features ...string) bool {
	if err == nil || len(features) <= 0 {
		return false
	}
	v := err.Error()
	return HasFeature(v, features...)
}
