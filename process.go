package main

import (
	"os"
)




var (
	maxtask,running,sleepin,stopped,zombied int
)


func page2k(n int)*Storage{
	storage := &Storage{}
	pagesize := os.Getpagesize()
	storage.Size = int64(n) * int64(pagesize)
	return storage
}

type procTab []*Proct

// 从/proc/#/stat 读取的数据
type Proct struct {
	pid		int
        comm		string
	state		byte
        ppid		int
        pgrp		int
        session		int
        tty_nr		int
        tpgid		int
        flags		uint
        minflt		uint64
        cminflt		uint64
        majflt		uint64
        cmajflt		uint64
        utime		uint64
        stime		uint64
        cutime		int64
        cstime		int64
        priority	int64
        nice		int64
        num_threads	int64
        itrealvalue	int64
        starttime	uint64
        vsize		uint64
        rss		int64
        rsslim		uint64
        startcode	uint64
        endcode		uint64
        startstack	uint64
        kstkesp		uint64
        kstkeip		uint64
        signal		uint64
        blocked		uint64
        sigignore	uint64
        sigcatch	uint64
        wchan		uint64
        nswap		uint64
        cnswap		uint64
        exit_signal	int
        processor	int
        rt_priority	uint
        policy		uint
        delayacct_blkio_ticks	uint64
        guest_time	uint64
        cguest_time	int64
        start_data	uint64
        end_data	uint64
        start_brk	uint64
        arg_start	uint64
        arg_end		uint64
        env_start	uint64
        env_end		uint64
	exit_code	int

	size		int     //   /proc/[pid]/statm
	resident	int	//   /proc/[pid]/statm
	share		int     //   /proc/[pid]/statm
	text		int     //   /proc/[pid]/statm
	data		int     //   /proc/[pid]/statm
}

func (p *Proct)ReadProc(pid string){
	pstat(pid,p)
	pstatm(pid,p)
}

func NewProcTab() procTab{
	pidl := pidList()
	ps := make([]*Proct,len(pidl))
	return  ps
}

func (pt procTab)Refresh(){
	maxtask,running,sleepin,stopped,zombied = 0,0,0,0,0

	pidl := pidList()
	for index,pid := range pidl{
		pt[index] = handleProc(pid)
	}
}
 
func handleProc(pid string)*Proct{
	proc := &Proct{}
	proc.ReadProc(pid)

	switch proc.state{
	case 'R':
		running += 1
	case 't':
		fallthrough
	case 'T':
		stopped += 1
	case 'Z':
		zombied += 1
	default:
		/* currently: 'D' (disk sleep),
               'I' (idle),
               'P' (parked),
               'S' (sleeping),
               'X' (dead - actually 'dying' & probably never seen)
 		*/
		sleepin += 1
	}
	maxtask += 1
	return proc
}
