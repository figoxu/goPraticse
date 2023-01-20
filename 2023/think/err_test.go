package think_test

import (
	"errors"
	"figoxu/think"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("error_is_as_of_official_lib_after_1.13", func() {
	var ErrTest = errors.New("error1")
	err1 := fmt.Errorf("return err %w", think.TestError{Value: "error1"})
	err2 := fmt.Errorf("inner error is %w", ErrTest)

	It("InfAnyNilMustNil", func() {
		Ω(errors.As(err1, &think.TestError{})).Should(BeTrue())
		Ω(errors.Is(err1, ErrTest)).Should(BeFalse())
		Ω(errors.As(err2, &think.TestError{})).Should(BeFalse())
		Ω(errors.Is(err2, ErrTest)).Should(BeTrue())
	})
})
