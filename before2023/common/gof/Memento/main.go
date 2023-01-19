package main

import (
	"fmt"
)

type GameRole struct {
	vit int
	atk int
	def int
}

func (this *GameRole) StateDisplay() {
	fmt.Println("角色当前状态：")
	fmt.Println("体力：", this.vit)
	fmt.Println("攻击力：", this.atk)
	fmt.Println("防御：", this.def)
	fmt.Println("")
}

func (this *GameRole) GetInitState() {
	this.vit = 100
	this.def = 100
	this.atk = 100
}

func (this *GameRole) Fight() {
	this.vit = 0
	this.def = 0
	this.atk = 0
}

func (this *GameRole) SaveState() RoleStateMemento {
	return RoleStateMemento{this.vit, this.atk, this.def}
}

func (this *GameRole) RecoveryState(memento RoleStateMemento) {
	this.vit = memento.vit
	this.atk = memento.atk
	this.def = memento.def
}

type RoleStateMemento struct {
	vit int
	atk int
	def int
}

type RoleStateCaretaker struct {
	memento RoleStateMemento
}

func main() {
	lixiaoyao := new(GameRole)
	lixiaoyao.GetInitState()
	lixiaoyao.StateDisplay()

	stateAdmin := new(RoleStateCaretaker)
	stateAdmin.memento = lixiaoyao.SaveState()

	lixiaoyao.Fight()
	lixiaoyao.StateDisplay()

	lixiaoyao.RecoveryState(stateAdmin.memento)
	lixiaoyao.StateDisplay()
}
