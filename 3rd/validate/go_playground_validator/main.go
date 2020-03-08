package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/quexer/utee"
	"github.com/sirupsen/logrus"
)

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"`
	Awesome        string     `validate:"is-awesome"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func main() {
	validate := validator.New()
	err := validate.RegisterValidation("is-awesome", ValidateAweSome)
	utee.Chk(err)
	logrus.Println("hello")

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	v := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Awesome:        "awesome",
		Addresses:      []*Address{address},
	}
	err = validate.Struct(v)
	utee.Chk(err)
}

func ValidateAweSome(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}
