package main

type CodeStack []Code

func (p CodeStack) Empty() bool        { return len(p) == 0 }
func (p CodeStack) Peek() Code  { return p[len(p)-1] }
func (p *CodeStack) Put(i Code) { (*p) = append((*p), i) }
func (p *CodeStack) Pop() Code {
	d := (*p)[len(*p)-1]
	(*p) = (*p)[:len(*p)-1]
	return d
}

func AppendCode(stacks CodeStack,val string)CodeStack{
	var code Code = Code(val)
	stacks = append(stacks,code)
	return stacks
}