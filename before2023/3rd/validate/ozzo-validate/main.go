package main

import (
	"fmt"
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

type Customer struct {
	Name    string
	Gender  string
	Email   string
	Address Address
}

func (a Address) Validate() error {
	return validation.StructRules{}.
		// Street cannot be empty, and the length must between 5 and 50
		Add("Street", validation.Required, validation.Length(5, 50)).
		// City cannot be empty, and the length must between 5 and 50
		Add("City", validation.Required, validation.Length(5, 50)).
		// State cannot be empty, and must be a string consisting of two letters in upper case
		Add("State", validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))).
		// Zip cannot be empty, and must be a string consisting of five digits
		Add("Zip", validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))).
		// performs validation
		Validate(a)
}

func (c Customer) Validate() error {
	return validation.StructRules{}.
		// Name cannot be empty, and the length must be between 5 and 20.
		Add("Name", validation.Required, validation.Length(5, 20)).
		// Gender is optional, and should be either "Female" or "Male".
		Add("Gender", validation.In("Female", "Male")).
		// Email cannot be empty and should be in a valid email format.
		Add("Email", validation.Required, is.Email).
		// Validate Address using its own validation rules
		Add("Address").
		// performs validation
		Validate(c)
}

func main() {
	c := Customer{
		Name:  "Qiang Xue",
		Email: "q",
		Address: Address{
			Street: "123 Main Street",
			City:   "Unknown",
			State:  "Virginia",
			Zip:    "12345",
		},
	}

	err := validation.Validate(c) // or alternatively, err := c.Validate()
	fmt.Println(err)
	// Output:
	// Address: (State: must be in a valid format.); Email: must be a valid email address.
}
