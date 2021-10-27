package main

import (
	"time"
)

var interval = 1 * time.Second
var gticker chan struct{}
var ticker *time.Ticker


func init(){
	gticker = make(chan struct{},1)
	ticker = time.NewTicker(interval)
}

func main(){
	ch := make(chan *Cell)
	go run(ch)
	printer(ch)
}

func run(ch chan<- *Cell){
	gticker <- struct{}{}
	for {
		select {
		case <- gticker:
			sender(ch)
			cpuSender(ch)
		case <- ticker.C:
			sender(ch)
			cpuSender(ch)
		}
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
