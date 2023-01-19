package main

import (
	"bytes"
	"compress/flate"
	"github.com/figoxu/utee"
	"io"
	"log"
)

func main() {
	b, err := deflate([]byte("hello world"))
	utee.Chk(err)
	log.Println(string(b))
	rb, err := inflate(b)
	utee.Chk(err)
	log.Println(string(rb))
}

// Compress with DEFLATE
func deflate(input []byte) ([]byte, error) {
	output := new(bytes.Buffer)
	if writer, err := flate.NewWriter(output, 1); err != nil {
		return nil, err
	} else if _, err = io.Copy(writer, bytes.NewBuffer(input)); err != nil {
		return nil, err
	} else if err = writer.Close(); err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

// Decompress with DEFLATE
func inflate(input []byte) ([]byte, error) {
	output := new(bytes.Buffer)
	reader := flate.NewReader(bytes.NewBuffer(input))
	if _, err := io.Copy(output, reader); err != nil {
		return nil, err
	} else if err = reader.Close(); err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}
