package main

import (
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

func readAll(filename string)[]byte{
	fi,err := os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer fi.Close()
	data, err := ioutil.ReadAll(fi)
	if err != nil{
		panic(err)
	}
	return data
}

func uptime()string{
	st := string(readAll("/proc/uptime"))
	st = strings.Split(st," ")[0]
	uptime,err := strconv.ParseFloat(st,32)
	if err !=nil{
		panic(err)
	}
	return FormatTime(float32(uptime))
}


func userCount()string{
	count := 0
	utmp := readAll("/var/run/utmp")
	for i:=0;i<len(utmp);i+=384{
                if utmp[i] == 7{
			count ++
		}
        }
	return fmt.Sprintf("  %d users",count)
}

func loadavg()string{
	sload := string(readAll("/proc/loadavg"))
	lload := strings.Split(sload," ")

	return fmt.Sprintf("  load average: %s, %s, %s",lload[0],lload[1],lload[2])
}

func pidList()[]string {
	lpid := make([]string,0)
	dir_list, err := ioutil.ReadDir("/proc")
        if err !=nil{
                panic("read dir error")
        }
	for _,v := range dir_list {
                name := v.Name()
                if name[0] >= '1' && name[0] <= '9'{
			lpid = append(lpid,name)
                }
        }
	return lpid

}

func pstat(pid string,p *Proct){
	stat := string(readAll("/proc/"+pid+"/stat"))
	fmt.Sscanf(stat,
		"%d " +			//pid
		"%s " +		//comm
		"%c " +			//state
		"%d " +			//ppid
		"%d " +			//pgrp
		"%d " +			//session
		"%d " +			//tty_nr
		"%d " +			//tpgid
		"%d " +			//flags
		"%d " +		//minflt
		"%d " +		//cminflt
		"%d " +		//majflt
		"%d " +		//cmajflt
		"%d " +		//utime
		"%d " +		//stime
		"%d " +		//cutime
		"%d " +		//cstime
		"%d " +		//priority
		"%d " +		//nice
		"%d " +		//num_threads
		"%d " +		//itrealvalue
		"%d " +		//starttime
		"%d " +		//vsize
		"%d " +		//rss
		"%v " +		//rsslim
		"%d " +		//startcode
		"%d " +		//endcode
		"%d " +		//startstack
		"%d " +		//kstkesp
		"%d " +		//kstkeip
		"%d " +		//signal
		"%d " +		//blocked
		"%d " +		//sigignore
		"%d " +		//sigcatch
		"%d " +		//wchan
		"%d " +		//nswap
		"%d " +		//cnswap
		"%d " +			//exit_signal
		"%d " +			//processor
		"%d " +			//rt_priority
		"%d " +			//policy
		"%d " +		//delayacct_blkio_ticks
		"%d " +		//guest_time
		"%d " +		//cguest_time
		"%d " +		//start_data
		"%d " +		//end_data
		"%d " +		//start_brk
		"%d " +		//arg_start
		"%d " +		//arg_end
		"%d " +		//env_start
		"%d " +		//env_end
		"%d",
		&p.pid,
		&p.comm,
		&p.state,
		&p.ppid,
		&p.pgrp,
		&p.session,
		&p.tty_nr,
		&p.tpgid,
		&p.flags,
		&p.minflt,
		&p.cminflt,
		&p.majflt,
		&p.cmajflt,
		&p.utime,
		&p.stime,
		&p.cutime,
		&p.cstime,
		&p.priority,
		&p.nice,
		&p.num_threads,
		&p.itrealvalue,
		&p.starttime,
		&p.vsize,
		&p.rss,
		&p.rsslim,
		&p.startcode,
		&p.endcode,
		&p.startstack,
		&p.kstkesp,
		&p.kstkeip,
		&p.signal,
		&p.blocked,
		&p.sigignore,
		&p.sigcatch,
		&p.wchan,
		&p.nswap,
		&p.cnswap,
		&p.exit_signal,
		&p.processor,
		&p.rt_priority,
		&p.policy,
		&p.delayacct_blkio_ticks,
		&p.guest_time,
		&p.cguest_time,
		&p.start_data,
		&p.end_data,
		&p.start_brk,
		&p.arg_start,
		&p.arg_end,
		&p.env_start,
		&p.env_end,
		&p.exit_code)
}

func pstatm(pid string, p *Proct){
	statm := string(readAll("/proc/"+pid+"/statm"))
	fmt.Sscanf(statm,"%d %d %d %d 0 %d 0",
		&p.size,&p.resident,&p.share,&p.text,&p.data)
}

func onlineCPU()[]string{
	cpus := make([]string,0)
	onCPUs := string(readAll("/sys/devices/system/cpu/online"))
	onCPUs = strings.TrimSpace(onCPUs)
	lonCPUs := strings.Split(onCPUs,",")
	for _,group := range lonCPUs{
		se := strings.Split(group, "-")
		start,err := strconv.Atoi(se[0])
		if err !=nil{
			panic(err)
		}
		end,err := strconv.Atoi(se[1])
		if err !=nil{
			panic(err)
		}
		for i := start; i<=end;i++{
			cpus = append(cpus, fmt.Sprintf("cpu%d",i))
		}
	}
	return cpus
}

func stat()[]string{
	stat := string(readAll("/proc/stat"))
	stat = strings.Replace(stat,"  "," ",-1)
	return strings.Split(stat,"\n")
}

func statCPUTime()map[string][]int64{
	ctimeStat := make(map[string][]int64)
	st := stat()
	onCpuCount := len(onlineCPU())
	for i := 0; i <= onCpuCount; i++ {
		cl := strings.Split(st[i]," ")
		ctimes := make([]int64,0)
		for _, t := range cl[1:]{
			itime,err := strconv.ParseInt(t, 10, 64)
			if err !=nil{
				panic(err)
			}
			ctimes = append(ctimes, itime)
		}
		ctimeStat[cl[0]] = ctimes
	}
	return ctimeStat
}

