package main

import (
	"time"
)

func main(){
	ch := make(chan *Cell)
	go sender(ch)
	printer(ch)
}

func sender(ch chan<- *Cell){
	for ; ;{
		cell := newCell()
		cell.msg = head()
		ch<- cell


		cell = newCell()
		cell.msg = tasks()
		cell.y = 1
		ch <- cell
		time.Sleep(1 * time.Second)
	}
}
