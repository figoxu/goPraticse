package b

import (
	"github.com/murlokswarm/errors"
	"github.com/quexer/utee"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

func InnerPanic() {
	defer Catch()
	utee.Chk(errors.New("测试错误"))
}

func Catch(hooks ...func(err ...interface{})) {
	if err := recover(); err != nil {
		logrus.WithFields(logrus.Fields{
			"stack":   string(debug.Stack()),
			"recover": true,
		}).Println("panic : ", err)
		for _, hook := range hooks {
			hook(err)
		}
	}
}
