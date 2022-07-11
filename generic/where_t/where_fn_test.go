package where_t_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/figoxu/where_t"
)

type Student struct {
	Id      int
	Name    string
	Age     int
	Score   int
	Gender  int
	ClassId int
}

type StudentWhere struct {
}

func (p *StudentWhere) IsAge(age int) where_t.WhereT[Student] {
	return func(x *Student) bool {
		return x.Age == age
	}
}

func (p *StudentWhere) IsName(name string) where_t.WhereT[Student] {
	return func(x *Student) bool {
		return x.Name == name
	}
}

func (p *StudentWhere) IsGender(gender int) where_t.WhereT[Student] {
	return func(x *Student) bool {
		return x.Gender == gender
	}
}

func (p *StudentWhere) IsClassId(classId int) where_t.WhereT[Student] {
	return func(x *Student) bool {
		return x.ClassId == classId
	}
}

func genStudents(count int) []*Student {
	genders := []int{0, 1}
	classIds := []int{1, 2, 3, 4, 5, 6, 7, 8}
	scores := []int{60, 66, 88, 80, 55, 37, 73, 93}
	names := []string{"andy", "jack", "lucy", "lily", "tom", "jason"}
	ages := []int{16, 17, 18, 19, 20, 21, 15, 14, 13}
	var out []*Student
	for i := 0; i < count; i++ {
		s := &Student{
			Id:      i,
			Name:    names[i%len(names)],
			Age:     ages[i%len(ages)],
			Score:   scores[i%len(scores)],
			Gender:  genders[i%len(genders)],
			ClassId: classIds[i%len(classIds)],
		}
		out = append(out, s)
	}
	return out
}

func TestX(t *testing.T) {
	j := func(in []*Student) string {
		b, _ := json.Marshal(in)
		return string(b)
	}
	vs := genStudents(1000)
	fmt.Println("BEFORE: ", j(vs))

	w := StudentWhere{}
	out := w.IsAge(16).And(
		w.IsClassId(1),
		w.IsName("andy"),
	).Val(vs...)

	fmt.Println("AFTER: ", j(out))
}
