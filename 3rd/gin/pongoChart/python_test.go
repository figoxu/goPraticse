package main

import (
	"testing"
	"fmt"
)

func Test_getTableNames(t *testing.T) {
	names := getTableNames("postgres", "rails_tpl", "figo", "123456", "localhost", "5432")
	for _, name := range names {
		fmt.Println(name)
	}
}
