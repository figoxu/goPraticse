package main
import (
	"net/http"
	"html/template"
	"github.com/figoxu/utee"
)



type Person struct {
	Name    string
	Age     int
	Emails  []string
	Company string
	Role    string
}

type OnlineUser struct {
	User      []*Person
	LoginTime string
}


func main(){
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8888", nil)
}


func Handler(w http.ResponseWriter, r *http.Request) {
	dumx := Person{
		Name: "zoro",
		Age: 27,
		Emails: []string{"dg@gmail.com", "dk@hotmail.com"},
		Company: "Omron",
		Role: "SE"}

	chxd := Person{Name: "chxd", Age: 27, Emails: []string{"test@gmail.com", "d@hotmail.com"}}

	onlineUser := OnlineUser{User: []*Person{&dumx, &chxd}}

	//t := template.New("Person template")
	//t, err := t.Parse(templ)
	t, err := template.ParseFiles("/home/figo/develop/env/GOPATH/src/github.com/figoxu/goPraticse/template/tmpl.html")
	utee.Chk(err)

	err = t.Execute(w, onlineUser)
	utee.Chk(err)
}
