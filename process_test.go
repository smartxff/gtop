package main

import (
	"testing"
)

func TestRefresh(t *testing.T){
	pt := NewProcTab()
	pt.Refresh()
	t.Log(maxtask,running,sleepin,stopped,zombied)
}
