package main

import (
	"github.com/bxcodec/faker/v4"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"os"
)

func main() {
	writer := &MockWriter{
		FileName: "/Users/xujianhui/mobvista/mtg/github/goPraticse/2023/bson/data.bson",
	}
	err := writer.Write(30)
	if err != nil {
		panic(err)
	}
}

type MockWriter struct {
	FileName string
}

func (p *MockWriter) Write(count int) error {
	for i := 0; i < count; i++ {
		vs, err := p.provider(p.mock())
		if err != nil {
			return err
		}
		err = p.appendBson(p.FileName, vs)
		if err != nil {
			return err
		}
	}
	return nil
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *MockWriter) mock() *Person {
	return &Person{
		Name: faker.ChineseName(),
		Age:  rand.Intn(18) + 18,
	}
}

func (p *MockWriter) provider(person *Person) ([]byte, error) {
	data, err := bson.Marshal(person)
	return data, errors.WithStack(err)
}

func (p *MockWriter) appendBson(fileName string, vs []byte) error {
	file, err := os.OpenFile("data.bson", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()
	_, err = file.Write(vs)
	return errors.WithStack(err)
}
