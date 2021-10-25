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
		cell.msg = top()
		ch<- cell
		time.Sleep(3 * time.Second)
	}
}
