package think_test

import (
	"errors"
	"figoxu/think"
	"fmt"
	"testing"
)

var ErrTest = errors.New("error1")

func returnError1() error {
	return fmt.Errorf("return err %w", think.TestError{Value: "error1"})
}

func returnError2() error {
	return fmt.Errorf("inner error is %w", ErrTest)
}

func TestErrorCheck(t *testing.T) {
	err1 := returnError1()
	err2 := returnError2()
	fmt.Println("Error.as", errors.As(err1, &think.TestError{}))
	fmt.Println("Error.is", errors.Is(err1, ErrTest))
	fmt.Println("Error.as", errors.As(err2, &think.TestError{}))
	fmt.Println("Error.is", errors.Is(err2, ErrTest))
}
