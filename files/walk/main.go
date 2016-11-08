
package main

// Importing the required packages
import (
	"fmt"
	"os"
	"path/filepath"
	"log"
)

// Main function
func main() {
	// Walk the directory tree using Walk() Golang method
	dir := "/home/figo/develop/env/GOPATH/src/github.com/figoxu/"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			log.Println("")
		}else {
			fmt.Println(path)
		}
		return nil
	})
}