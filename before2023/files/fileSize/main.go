package main
import (
	"path/filepath"
	"os"
	"log"
)

func main(){
	if size ,err := DirSize("../../..");err!=nil {
		log.Println("@err:",err)
	}else{
		log.Println("@size:", size%1024 ," KB")
	}


}
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}