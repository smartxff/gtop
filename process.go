package main



const (
	PID = iota
	COMM
	STATE
	PPID
	PGRP
	SESSION
	TTY_NR
	TPGID
	MINFLT
	CMINFLT
	MAJFLT
	CMAJFLT
	UTIME
	STIME
	CUTIME
	CSTIME
	PRIORITY
	NICE
	NUM_THREADS
	ITREALVALUE
	STARTTIME
	RSS
	RSSLIM
	STARTCODE
	ENDCODE
	STARTSTACK
	KSTKESP
	KSTKEIP
	SIGNAL
	BLOCKED
	SIGIGNORE
	SIGCATCH
	WCHAN
	NSWAP
	CNSWAP
	EXIT_SIGNAL
	PROCESSOR
	RT_PRIORITY
	DELAYACCT_BLKIO_TICKS
	GUEST_TIME
	CGUEST_TIME
	START_DATA
	END_DATA
	START_BRK
	ARG_START
	ARG_END
	ENV_START
	ENV_END
	EXIT_CODE

)


var (
	maxtask,running,sleepin,stopped,zombied int
)

type procTab []*Proct

// 从/proc/#/stat 读取的数据
type Proct struct {
	state	string

}

func (p *Proct)ReadProc(pid string){
	stat := pstat(pid)
	p.state = stat[STATE]
}

func NewProcTab() *procTab{
	pt := new(procTab)
	return  pt
}

func (pt *procTab)Refresh(){
	maxtask,running,sleepin,stopped,zombied = 0,0,0,0,0
	pidl := pidList()
	ps := make([]*Proct,len(pidl))


	for index,pid := range pidl{
		ps[index] = handleProc(pid)
	}
	ppt := procTab(ps)
	pt = &ppt
}

func handleProc(pid string)*Proct{
	proc := &Proct{}
	proc.ReadProc(pid)

	switch proc.state{
	case "R":
		running += 1
	case "t":
		fallthrough
	case "T":
		stopped += 1
	case "Z":
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
