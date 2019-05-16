package ormsample

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/quexer/utee"
)

type IntArray []int

func NewIntArray(ints ...int) IntArray {
	return IntArray(ints)
}

func (p *IntArray) String() string {
	if len(*p) <= 0 {
		return "{}"
	}
	b, e := json.Marshal(p)
	utee.Chk(e)
	val := string(b)
	val = fmt.Sprint("{", val[1:len(val)-1], "}")
	return val
}

func (p *IntArray) Scan(val interface{}) error {
	b := val.([]byte)
	v := string(b)
	v = fmt.Sprint("[", v[1:len(v)-1], "]")
	return json.Unmarshal([]byte(v), p)
}

func (p IntArray) Value() (driver.Value, error) {
	return p.String(), nil
}
