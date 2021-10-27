package main

import (
	"fmt"
	"os"
	"github.com/nsf/termbox-go"
	"github.com/mattn/go-runewidth"
)

const (
	DEFAULTOP	= iota
	FLUSH
	CLEAR
)


type Cell struct {
	x    int
	y    int
	msg   string
	fg   termbox.Attribute
	bg   termbox.Attribute
	op int
}


func newCell()*Cell{
	cell := &Cell{}
        cell.x = 0
        cell.y = 0
        cell.msg = "top - "
        cell.fg = termbox.ColorDefault
        cell.bg = termbox.ColorDefault
	cell.op = DEFAULTOP
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
		switch cell.op{
			case FLUSH:
				termbox.Flush()
			case CLEAR:
				termbox.Clear(termbox.ColorDefault,termbox.ColorDefault)

		}
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

func run(ch chan<- *Cell){
        gticker <- struct{}{}
        for {
		clear(ch)
                select {
                case <- gticker:
                        sender(ch)
                        cpuSender(ch)
                case <- ticker.C:
                        sender(ch)
                        cpuSender(ch)
                }
                flush(ch)
        }
}
func sender(ch chan<- *Cell){
        cell := newCell()
        cell.msg = head()
        ch<- cell


        cell = newCell()
        cell.msg = tasks()
        cell.y = 1
        ch <- cell
}

func flush(ch chan<- *Cell){
        cell := newCell()
        cell.op = FLUSH
        ch <- cell
}

func clear(ch chan<- *Cell){
        cell := newCell()
        cell.op = CLEAR
        ch <- cell
}

