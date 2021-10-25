package main

import (
	"testing"
)

func TestFormatTime(t *testing.T){
	t.Log(FormatTime(56))
	t.Log(FormatTime(100))
	t.Log(FormatTime(4000))
	t.Log(FormatTime(3091243.29))
} 
