package main

import (
	"github.com/nsf/termbox-go"
)

func main(){
	ch := make(chan *Cell)
	go genCell(ch)
	Print(ch)
}

func genCell(ch chan<- *Cell){
	cell := &Cell{}
	cell.x = 0
	cell.y = 0
	cell.msg = "top - "
	cell.fg = termbox.ColorDefault
	cell.bg = termbox.ColorDefault
	ch<- cell
}
