package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
)

type Box [4][4]int

func (t *Box) MergeAndReturnKey() termbox.Key {
	var changed bool
Lable:
	changed = false
	ev := termbox.PollEvent()
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			changed = t.MergeUP()
		case termbox.KeyArrowDown:
			changed = t.MergeDwon()
		case termbox.KeyArrowLeft:
			changed = t.MergeLeft()
		case termbox.KeyArrowRight:
			changed = t.MergeRight()
		case termbox.KeyEsc, termbox.KeyEnter:
			changed = true
		default:
			changed = false
			//t.Print(0, 3)
		}
		if !changed {
			goto Lable
		}

	case termbox.EventResize:
		x, y := termbox.Size()
		t.Print(x/2-10, y/2-4)
		goto Lable
	case termbox.EventError:
		panic(ev.Err)
	}
	step++
	return ev.Key
}

func (t *Box) MergeUP() bool {
	tl := len(t)
	changed := false
	notfull := false
	for i := 0; i < tl; i++ {

		np := tl //the last number needed check, first time use len(t).
		n := 0   //count of none 0.

		//clean 0 from top to the last number and move numbers together.
		//imag another t that smaller than this, but covered this.
		//n after "for" is size of the small t,  gives the value of next np.
		for x := 0; x < np; x++ {
			if t[x][i] != 0 {
				t[n][i] = t[x][i]
				if n != x {
					changed = true
				}
				n++
			}
		}
		if n < tl {
			notfull = true
		}
		np = n
		//mergeup all the number x that are same with its uper one.
		//uper one store 2*x, downer store 0.
		for x := 0; x < np-1; x++ {
			if t[x][i] == t[x+1][i] {
				t[x][i] *= 2
				t[x+1][i] = 0
				Score += t[x][i] * step
				x++
				changed = true
				//	n--
			}
		}
		//clean the new added 0 use the same way.
		n = 0
		for x := 0; x < np; x++ {
			if t[x][i] != 0 {
				t[n][i] = t[x][i]
				n++
			}
		}
		//cover the unchecked with 0
		for x := n; x < tl; x++ {
			t[x][i] = 0
		}
	}
	return changed || !notfull
}
func (t *Box) MergeDwon() bool {
	t.MirrorV()
	changed := t.MergeUP()
	t.MirrorV()
	return changed
}
func (t *Box) MergeLeft() bool {
	t.Right90()
	changed := t.MergeUP()
	t.Left90()
	return changed
}
func (t *Box) MergeRight() bool {
	t.Left90()
	changed := t.MergeUP()
	t.Right90()
	return changed
}

func (t Box) Print(ox, oy int) error {
	fg := termbox.ColorYellow
	bg := termbox.ColorBlack
	termbox.Clear(fg, bg)
	str := "  SCORE: " + fmt.Sprint(Score)
	for n, c := range str {
		termbox.SetCell(ox+n, oy-1, c, fg, bg)
	}
	str = "ESC:exit " + "Enter:replay"
	for n, c := range str {
		termbox.SetCell(ox+n, oy-2, c, fg, bg)
	}
	str = "Play with arrow key!"
	for n, c := range str {
		termbox.SetCell(ox+n, oy-3, c, fg, bg)
	}
	fg = termbox.ColorBlack
	bg = termbox.ColorGreen
	for i := 0; i <= len(t); i++ {
		for x := 0; x < 5*len(t); x++ {
			termbox.SetCell(ox+x, oy+i*2, '-', fg, bg)
		}
		for x := 0; x <= 2*len(t); x++ {
			if x%2 == 0 {
				termbox.SetCell(ox+i*5, oy+x, '+', fg, bg)
			} else {
				termbox.SetCell(ox+i*5, oy+x, '|', fg, bg)
			}
		}
	}
	fg = termbox.ColorYellow
	bg = termbox.ColorBlack
	for i := range t {
		for j := range t[i] {
			if t[i][j] > 0 {
				str := fmt.Sprint(t[i][j])
				for n, char := range str {
					termbox.SetCell(ox+j*5+1+n, oy+i*2+1, char, fg, bg)
				}
			}
		}
	}
	return termbox.Flush()
}

func (t *Box) Transpose() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[j][i] = num
		}
	}
	*t = *tn
}
func (t *Box) Right90() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[j][len(t)-i-1] = num
		}
	}
	*t = *tn
}
func (t *Box) Left90() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[len(line)-j-1][i] = num
		}
	}
	*t = *tn
}
func (t *Box) MirrorV() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[len(t)-i-1][j] = num
		}
	}
	*t = *tn
}
func (t *Box) MirrorH() {
	tn := new(Box)
	for i, line := range t {
		for j, num := range line {
			tn[i][len(line)-j-1] = num
		}
	}
	*t = *tn
}
func (t *Box) CheckWinAndAdd() Status {
	for _, x := range t {
		for _, y := range x {
			if y >= Max {
				return Win
			}
		}
	}
	i := rand.Intn(len(t))
	j := rand.Intn(len(t))
	for x := 0; x < len(t); x++ {
		for y := 0; y < len(t); y++ {
			if t[i%len(t)][j%len(t)] == 0 {
				t[i%len(t)][j%len(t)] = 2 << (rand.Uint32() % 2)
				return Add
			}
			j++
		}
		i++
	}
	return Lose
}
