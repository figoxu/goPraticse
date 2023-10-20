package main

import (
	"bufio"
	//"bytes"
	"encoding/json"
	"fmt"
	"github.com/bxcodec/faker/v4"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"

	"math/rand"
	"os"
)

func main() {
	fileName := "/Users/xujianhui/mobvista/mtg/github/goPraticse/2023/bson/data.bson"
	writer := &MockWriter{
		FileName: fileName,
	}
	err := writer.Write(30 * 10000)
	if err != nil {
		panic(err)
	}
	ch := make(chan Person)
	reader := &BufferReader{
		FileName:  fileName,
		ChunkSize: 1024,
	}
	go func() {
		reader.read(ch)
	}()
	i := 0
	for v := range ch {
		bs, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bs))
		i++
	}
	fmt.Println("total is ", i)
}

type BufferReader struct {
	FileName  string
	ChunkSize int
}

func (p *BufferReader) read(out chan Person) error {
	file, err := os.Open(p.FileName)
	if err != nil {
		return errors.WithStack(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder, err := bson.NewDecoder(reader)

	// 定义每个分段的大小
	chunkSize := 10 * 1024 * 1024 // 10 MB

	// 分段读取BSON文档
	var docs []bson.M
	for {
		// 读取下一个分段
		chunk := make([]byte, chunkSize)
		_, err := file.Read(chunk)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				panic(err)
			}
		}

		// 解码分段中的所有文档
		for len(chunk) > 0 {
			var doc bson.M
			remaining, err := decoder.DecodeBytes(chunk, &doc)
			if err != nil {
				panic(err)
			}

			// 将解码完的文档添加到列表中
			docs = append(docs, doc)

			// 更新剩余的分段数据
			chunk = remaining
		}
	}

	return nil
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
	Name string `bson:"name"`
	Age  int    `bson:"age"`
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
