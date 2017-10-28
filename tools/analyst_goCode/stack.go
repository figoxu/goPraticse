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

func (p CodeStack) split(braceLeft,braceRight string) CodeStack{
	var preStack,codeStack CodeStack = make([]Code,0),make([]Code,0)
	for _,code:=range p {
		if code=="" {
			continue
		}
		codes:=code.Split(braceLeft,true)
		preStack = append(preStack,codes...)
	}
	for _,code:=range preStack{
		if code=="" {
			continue
		}
		if string(code)==braceLeft {
			codeStack = append(codeStack,code)
			continue
		}
		codes:=code.Split(braceRight,false)
		codeStack = append(codeStack,codes...)
	}
	return codeStack
}

