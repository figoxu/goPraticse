package ut

import (
	"fmt"

	"github.com/pkg/errors"
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

func Recovery() {
	if err := recover(); err != nil {
		msg := fmt.Sprintf("%+v", err)
		if v, ok := err.(error); ok && v != nil {
			msg = fmt.Sprintf("%+v", errors.WithStack(v))
		}
		logrus.WithField(`error`, msg).Println(`panic for recovery`)
	}
}
