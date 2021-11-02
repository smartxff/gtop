package main

import (
	"strings"
	"fmt"
)

const (
        CPU_USER = iota
        CPU_NICE
        CPU_SYSTEM
        CPU_IDLE
        CPU_IOWAIT
        CPU_IRQ
        CPU_SOFTIRQ
        CPU_STEAL
        CPU_GUEST
        CPU_GUEST_NICE
)

var cpuDetail bool

var cpuStat map[string][]int64

func init(){
	cpuDetail = false
	cpuStat = statCPUTime()
}



func cpus()map[string][]float32{
	cpuUtil := make(map[string][]float32)

	cpustime2 := statCPUTime()
	for icpu,times := range cpustime2 {
		c1 := cpuStat[icpu]
		c2 := cpustime2[icpu]
		util := make([]float32,len(times))
		var totle float32 = 0
		for index,_ := range times {
			util[index] = float32(c2[index] - c1[index])
			totle += util[index]
		}
		for i,u := range util {
			if totle == 0 {
				util[i] = 0
			}else {
				util[i] = u/totle * 100
			}
		}
		cpuUtil[icpu] = util
	}
	cpuStat = cpustime2
	return cpuUtil
}
func cpusFormat(cpus map[string][]float32)[]string{
	smaps := make([]string, 0)
	onlineCpus := onlineCPU()
	onlineCpus = append([]string{"cpu"},onlineCpus...)
	for _,cindex := range onlineCpus {
		times := cpus[cindex]
		s := fmt.Sprintf("%%%-6s:%5.1f us,%5.1f sy,%5.1f ni,%5.1f id,%5.1f wa,%5.1f hi,%5.1f si,%5.1f st",
			replaceCpuName(cindex),times[CPU_USER],times[CPU_SYSTEM],times[CPU_NICE],times[CPU_IDLE],
			times[CPU_IOWAIT],times[CPU_IRQ],times[CPU_SOFTIRQ],times[CPU_STEAL])
		smaps = append(smaps, s)
	}
	return smaps
}


func replaceCpuName(cname string)string{
	if cname == "cpu"{
		return "Cpu(s)"
	}
	return strings.Replace(cname,"c","C",1)

}

func cpuSender(ch chan<- *Cell){
	y := 2
	cs := cpusFormat(cpus())
	if cpuDetail {
		for index,cpuUtil := range cs[1:]{
			cell := newCell()
			cell.msg = cpuUtil
			cell.y = y + index
			ch <- cell
		}
	}else {
		cell := newCell()
		cell.msg = cs[0]
		cell.y = y
		ch <- cell
	}
}
