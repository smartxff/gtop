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
