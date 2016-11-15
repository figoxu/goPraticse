package main

import (
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

	user := User{Name: "Jinzhu", Age: 18, Role: "Admin"}
	employee := Employee{}
	copier.Copy(&employee, &user)


	// employee => Employee{ Name: "Jinzhu",           // Copy from field
	//                       Age: 18,                  // Copy from field
	//                       DoubleAge: 36,            // Copy from method
	//                       EmployeeId: 0,            // Just ignored
	//                       SuperRule: "Super Admin", // Copy to method
	//                      }

	// Copy struct to slice
	user2 := User{Name: "hello", Age: 18, Role: "User"}
	employees := []Employee{}
	copier.Copy(&employees, &user2)
	// employees => [{hello 18 0 36 Super User}]

	// Copy slice to slice
	users := []User{{Name: "Jinzhu", Age: 18, Role: "Admin"}, {Name: "jinzhu 2", Age: 30, Role: "Dev"}}
	employees2 := []Employee{}
	copier.Copy(&employees2, &users)
}
