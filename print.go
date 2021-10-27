package main

import (
	"fmt"
	"os"
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
)

type Cell struct {
	x    int
	y    int
	msg   string
	fg   termbox.Attribute
	bg   termbox.Attribute
}


func newCell()*Cell{
	cell := &Cell{}
        cell.x = 0
        cell.y = 0
        cell.msg = "top - "
        cell.fg = termbox.ColorDefault
        cell.bg = termbox.ColorDefault
        return cell
}

func printer( ch <-chan *Cell){
	err := termbox.Init()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	termbox.SetInputMode(termbox.InputEsc)
        go handleEvent()
	for cell := range ch {
		tbprint(cell.x, cell.y, cell.fg, cell.bg, cell.msg)
		termbox.Flush()
	}
}


func tbprint(x, y int, fg, bg termbox.Attribute, msg string){
	for _, c := range msg { 
		termbox.SetCell(x, y,c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func handleEvent(){
        for {
                switch ev := termbox.PollEvent(); ev.Type {
                case termbox.EventKey:
                        if ev.Key == termbox.KeyCtrlC {
                                termbox.Close()
                                os.Exit(0)
                        }else if ev.Ch == '1' {
				cpuDetail = !cpuDetail
				gticker <- struct{}{}
			}

                case termbox.EventError:
                        panic(ev.Err)

                }
        }
}

