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
