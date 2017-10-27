package main

type Node interface {
	AddChild(node Node)
	RemoveChild(node Node)
	RemoveById(id string)
	GetParent() Node
	GetId() string
}

type BaseNode struct {
	Id       string
	Parent   *BaseNode `json:"-"`
	Children []*BaseNode
}

func (p *BaseNode) AddChild(node *BaseNode) {
	if p.Children == nil {
		p.Children = make([]*BaseNode, 0)
	}
	p.Children = append(p.Children, node)
	node.Parent = p
}

func (p *BaseNode) RemoveChild(node *BaseNode) {
	p.RemoveById(node.Id)
}

func (p *BaseNode) RemoveById(id string) {
	cs := make([]*BaseNode, 0)
	for _, child := range p.Children {
		if child.GetId() != id {
			cs = append(cs, child)
		} else {
			child.Parent = nil
		}
	}
	p.Children = cs
}

func (p *BaseNode) GetParent() *BaseNode {
	return p.Parent
}

func (p *BaseNode) GetId() string {
	return p.Id
}
