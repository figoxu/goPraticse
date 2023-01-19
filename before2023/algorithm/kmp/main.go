package main
import (
	"fmt"
	"github.com/paddie/gokmp"
)

const str = "aabaabaaaabbaabaabaaabbaabaabb"
//          "        _          _      _   "
//                   8          19     26
const pattern = "aabb"

func main() {
	kmp, _ := gokmp.NewKMP(pattern)
	ints := kmp.FindAllStringIndex(str)

	fmt.Println(ints)
}