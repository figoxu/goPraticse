package main

import (
	"bytes"
	"compress/gzip"
	"github.com/quexer/utee"
	"io/ioutil"
	"log"
)

func main() {
	b, err := gzipCompress([]byte("hello world"))
	utee.Chk(err)
	log.Println(string(b))
	rb, err := gunzip(b)
	utee.Chk(err)
	log.Println(string(rb))
}

// Compress with GZIP
func gzipCompress(input []byte) ([]byte, error) {
	var output bytes.Buffer
	writer := gzip.NewWriter(&output)
	writer.Write(input)
	err := writer.Close()
	return output.Bytes(), err
}

// Decompress with GUNZIP
func gunzip(input []byte) ([]byte, error) {
	in := bytes.NewReader(input)
	if reader, err := gzip.NewReader(in); err != nil {
		return nil, err
	} else if fileContents, err := ioutil.ReadAll(reader); err != nil {
		return nil, err
	} else if err = reader.Close(); err != nil {
		return nil, err
	} else {
		return fileContents, err
	}
}
