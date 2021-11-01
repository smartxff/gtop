package main

import (
	"testing"
)

func TestMeminfo(t *testing.T){
	minfo := NewMeminfo()
	for k,v := range minfo{
		t.Log(k,v)
	}
	t.Log(minfo.Totle().MB())
	t.Log(minfo.MemTotle().MB())
	t.Log(minfo.SwapTotle().MB())
	t.Log(minfo.MemFree().MB())
	t.Log(minfo.MemUsed().MB())
	t.Log(minfo.BufferCache().MB())
	t.Log(minfo.MIBMem())
	t.Log(minfo.MIBSwap())
}
