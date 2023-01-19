package ut

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

func (p *Lg) WithDebug(k string, v interface{}) *Lg {
	p.Lock()
	defer p.Unlock()
	p.m[k] = v
	return p
}

type Lg struct {
	sync.Mutex
	m  map[string]interface{}
	rs *logrus.Entry
}

func Log(entry ...*logrus.Entry) *Lg {
	rs := logrus.WithFields(logrus.Fields{})
	if len(entry) > 0 {
		rs = entry[0]
	}
	return &Lg{
		rs: rs,
		m:  map[string]interface{}{},
	}
}

func (p *Lg) WithErr(err error) *Lg {
	msg := fmt.Sprintf("%+v", err)
	p.rs = p.rs.WithField("err", msg)
	return p
}

func (p *Lg) debug() *Lg {
	if len(p.m) == 0 {
		return p
	}

	p.rs = p.rs.WithField("debug", JsonString(p.m))
	return p
}

func (p *Lg) Entry() *logrus.Entry {
	return p.debug().rs
}
