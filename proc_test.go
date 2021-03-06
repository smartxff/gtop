package main

import "testing"

func TestReadAll(t *testing.T){
	s := readAll("/proc/uptime")
	t.Log(s)
}

func TestUptime(t *testing.T){
	st := uptime()
	t.Log(st)
}

func TestTimeNow(t *testing.T){
	now := timeNow()
	t.Log(now)
}
func TestUserCount(t *testing.T){
	t.Log(userCount())
}

func TestPidList(t *testing.T){
	t.Log(pidList())
}
func TestPstat(t *testing.T){
	proc := &Proct{}
	pstat("1", proc)
	t.Log(proc)
}
func TestOnlineCPU(t *testing.T){
	t.Log(onlineCPU())
}
func TestStatCPUTime(t *testing.T){
	t.Log(statCPUTime())
}
