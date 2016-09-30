package is
import (
	"github.com/cheekybits/is"
	"testing"
	"github.com/figoxu/utee"
	"github.com/gogap/errors"
	"log"
)


func TestSomething(t *testing.T) {
	is := is.New(t)

	// ensure not nil
	is.OK(true)
	// ensure no error
	is.NoErr(nil)
//	is.OK(false)
	is.OK("hello")
	is.OK(len("hello"))
	is.OK(func(){
		utee.Chk(errors.New("err"))
	})

	// ensure many things in one go
	is.OK(true, errors.New("err2"), 13, "something")

	// ensure something does panic
	is.Panic(func(){
		log.Panic("cool")
	})
	is.PanicWith("tips over there", func(){
		log.Panic("err with tips")
	})

	// make sure two values are equal
	is.Equal(1, 2)

}