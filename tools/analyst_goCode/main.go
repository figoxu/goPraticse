package main

import (
	"fmt"
	"github.com/figoxu/Figo"
)

func main() {
	seq := Figo.NewSeqMem()
	seqStr := func() string {
		return fmt.Sprint(seq.Next())
	}
	tree := &BaseNode{
		Id: seqStr(),
	}
	for lv1Count := 0; lv1Count < 5; lv1Count++ {
		node := &BaseNode{
			Id: seqStr(),
		}
		tree.AddChild(node)
		for lv2 := 0; lv2 < 2; lv2++ {
			node.AddChild(&BaseNode{
				Id: seqStr(),
			})
		}
	}
	Figo.PrintJson("==>", tree)
	tree.RemoveChild(&BaseNode{Id: "2"})
	Figo.PrintJson("==>", tree)
}
