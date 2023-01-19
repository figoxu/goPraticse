package main

import (
	"fmt"
	"github.com/fatih/structtag"
	"reflect"
	"sort"
)

func main() {
	myDemo()
	fieldValues()
	sampleDemo()
}

func myDemo() {
	user := User{}
	user.Id = 12
	user.Name = "figo"
	fmt.Println(user.Base.Id)

	fmt.Println("%T", user)
	fmt.Println("%T", &user)
	fmt.Println("%T", &user)
	tp := reflect.TypeOf(user)
	fmt.Printf(" %v   %v \n", tp, tp.Kind())
	//tp = reflect.TypeOf(&user)
	//fmt.Printf(" %v   %v \n", tp, tp.Kind())
	idField, _ := tp.FieldByName("Base")
	sbtp := idField.Type
	fmt.Printf(" %v   %v \n", sbtp, sbtp.Kind())

	reflect.ValueOf(idField)
	//base:=idField.(*Base)

	val := reflect.ValueOf(&user).Elem()
	base := val.FieldByName("Base").Interface().(Base)
	fmt.Println(base.Id)
}

func fieldValues() {
	type Foo struct {
		FirstName string `tag_name:"tag 1"`
		LastName  string `tag_name:"tag 2"`
		Age       int    `tag_name:"tag 3"`
	}
	f := &Foo{
		FirstName: "Drew",
		LastName:  "Olson",
		Age:       30,
	}
	val := reflect.ValueOf(f).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Tag Value: %s\n", typeField.Name, valueField.Interface(), tag.Get("tag_name"))
	}
}

func sampleDemo() {
	type t struct {
		t string `json:"foo,omitempty,string" xml:"foo"`
	}

	// get field tag
	tag := reflect.TypeOf(t{}).Field(0).Tag

	// ... and start using structtag by parsing the tag
	tags, err := structtag.Parse(string(tag))
	if err != nil {
		panic(err)
	}

	// iterate over all tags
	for _, t := range tags.Tags() {
		fmt.Printf("tag: %+v\n", t)
	}

	// get a single tag
	jsonTag, err := tags.Get("json")
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonTag)         // Output: json:"foo,omitempty,string"
	fmt.Println(jsonTag.Key)     // Output: json
	fmt.Println(jsonTag.Name)    // Output: foo
	fmt.Println(jsonTag.Options) // Output: [omitempty string]

	// change existing tag
	jsonTag.Name = "foo_bar"
	jsonTag.Options = nil
	tags.Set(jsonTag)

	// add new tag
	tags.Set(&structtag.Tag{
		Key:     "hcl",
		Name:    "foo",
		Options: []string{"squash"},
	})

	// print the tags
	fmt.Println(tags) // Output: json:"foo_bar" xml:"foo" hcl:"foo,squash"

	// sort tags according to keys
	sort.Sort(tags)
	fmt.Println(tags) // Output: hcl:"foo,squash" json:"foo_bar" xml:"foo"
}

type Base struct {
	Id int
}

type User struct {
	Base
	Name string
}