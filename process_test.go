package main

import (
	"testing"
)

func TestRefresh(t *testing.T){
	pt := NewProcTab()
	pt.Refresh()
	t.Log(maxtask,running,sleepin,stopped,zombied)
	for _,p := range pt{
		if p.pid == 1{
			vsize := &Storage{Size: int64(p.vsize)}
			t.Log(p.pid,p.priority,p.nice,vsize.Kb(),page2k(int(p.rss)).Kb(),page2k(p.share).Kb(),p.state)
			t.Log(p.share,p.size,p.resident,p.text,p.data)
		}
	}
}

