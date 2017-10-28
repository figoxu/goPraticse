package main

import (
	"github.com/figoxu/Figo"
	"fmt"
)

var (
	seq=Figo.NewSeqMem()
	seqStr = func() string {
		return fmt.Sprint(seq.Next())
	}
)

func ParseStackToNode(codeStack CodeStack,braceLeft,braceRight string)CodeNode{
	rootNode := &CodeNode{
		Id:seqStr(),
		CodeStack: make(CodeStack, 0),
		Children:  make([]*CodeNode, 0),
	}
	codeStack=codeStack.split(braceLeft,braceRight)
	curNode,length:=rootNode,len(codeStack)
	for index,v:=range codeStack {
		if curNode.CodeStack==nil {
			curNode.CodeStack=make(CodeStack,0)
		}
		str:=string(v);
		if str==braceLeft {
			continue
		}
		nextStr:=""
		if index<length-1 {
			nextStr=string(codeStack[index+1])
		}
		if nextStr!=braceLeft && str!=braceRight {
			curNode.CodeStack=append(curNode.CodeStack, v)
		}else if nextStr==braceLeft {
			if curNode.Children==nil {
				curNode.Children = make([]*CodeNode,0)
			}
			node := &CodeNode{
				Id:        seqStr(),
				CodeStack: make(CodeStack, 0),
				Children:  make([]*CodeNode, 0),
			}
			node.Parent = curNode
			curNode.Children = append(curNode.Children,node)
			curNode = node
			curNode.CodeStack=append(curNode.CodeStack, v)
		}else if str==braceRight{
			curNode = curNode.Parent
		}
	}
	return *rootNode
}


