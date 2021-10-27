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
