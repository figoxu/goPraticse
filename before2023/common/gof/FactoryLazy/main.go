package main

import (
	"fmt"
)

type LeiFeng struct {
}

func (l *LeiFeng) sweep() {
	fmt.Println("扫地")
}

func (l *LeiFeng) wash() {
	fmt.Println("洗衣")
}

func (l *LeiFeng) buyRice() {
	fmt.Println("买米")
}

type Undergraduate struct {
	LeiFeng
}

type Volunteer struct {
	LeiFeng
}

type Ifactory interface {
	createLeiFeng() LeiFeng
}

type UndergraduateFactory struct {
}

func (u *UndergraduateFactory) createLeiFeng() LeiFeng {
	return new(Undergraduate).LeiFeng
}

type VolunteerFactory struct {
}

func (v *VolunteerFactory) createLeiFeng() LeiFeng {
	return new(Volunteer).LeiFeng
}

func main() {
	ifac := new(UndergraduateFactory)
	student := ifac.createLeiFeng()
	student.wash()
	student.sweep()
	student.buyRice()

}
