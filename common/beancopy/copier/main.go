package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Role string
	Age  int32
}

func (user *User) DoubleAge() int32 {
	return 2 * user.Age
}

type Employee struct {
	Name      string
	Age       int32
	DoubleAge int32
	EmployeId int64
	SuperRule string
}

func (employee *Employee) Role(role string) {
	employee.SuperRule = "Super " + role
}

func main() {
	var (
		user      = User{Name: "Jinzhu", Age: 18, Role: "Admin"}
		users     = []User{{Name: "Jinzhu", Age: 18, Role: "Admin"}, {Name: "jinzhu 2", Age: 30, Role: "Dev"}}
		employee  = Employee{}
		employees = []Employee{}
	)
	copier.Copy(&employee, &user)

	fmt.Printf("%#v \n", employee)
	copier.Copy(&employees, &user)

	fmt.Printf("%#v \n", employees)

	// Copy slice to slice
	employees = []Employee{}
	copier.Copy(&employees, &users)

	fmt.Printf("%#v \n", employees)
}