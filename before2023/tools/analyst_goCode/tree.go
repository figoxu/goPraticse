package main

type CodeNode struct {
	Id        string
	CodeStack CodeStack
	Parent    *CodeNode `json:"-"`
	Children  []*CodeNode
}

func (p *CodeNode) AddChild(node *CodeNode) {
	if p.Children == nil {
		p.Children = make([]*CodeNode, 0)
	}
	p.Children = append(p.Children, node)
	node.Parent = p
}

func (p *CodeNode) RemoveChild(node *CodeNode) {
	p.RemoveById(node.Id)
}

func (p *CodeNode) RemoveById(id string) {
	cs := make([]*CodeNode, 0)
	for _, child := range p.Children {
		if child.GetId() != id {
			cs = append(cs, child)
		} else {
			child.Parent = nil
		}
	}
	p.Children = cs
}

func (p *CodeNode) GetParent() *CodeNode {
	return p.Parent
}

func (p *CodeNode) GetId() string {
	return p.Id
}
