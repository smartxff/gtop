package main

import (
	"testing"
)

func TestCpu(t *testing.T){
	cs := cpus()
	t.Log(cpusFormat(cs))
}
