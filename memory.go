package main

import (
	"strings"
	"fmt"
	"regexp"
	"strconv"
)

const (
	KB = 1024
	MB = 1024 * 1024
	GB = 1024 * 1024
)

type Storage struct {
	Size int64
}

func (s *Storage)MB()string{
	return fmt.Sprintf("%.1f",float64(s.Size)/float64(MB))
}
func (s *Storage)Kb()int64{
	return s.Size / KB
}
func (s *Storage)Add(ss *Storage)*Storage{
	size := s.Size + ss.Size
	storage := &Storage{Size: size}
	return storage
}
func (s *Storage)Sub(ss *Storage)*Storage{
	size := s.Size - ss.Size
	storage := &Storage{Size: size}
	return storage
}

func (mi Meminfo)Find(name string) *Storage{
	m := map[string]*Storage(mi)
	return m[name]
}

type Meminfo map[string]*Storage


func (mi Meminfo)Totle()*Storage{
	totle := mi.MemTotle()
	return totle.Add(mi.SwapTotle())
}

func (mi Meminfo)MemTotle()*Storage{
	mtotle := mi.Find("MemTotal")
	return mtotle
}

func (mi Meminfo)MemFree()*Storage{
	free := mi.Find("MemFree")
	return free
}

func (mi Meminfo)Buffer()*Storage{
	buffer := mi.Find("Buffers")
	return buffer
}

func (mi Meminfo)Cache()*Storage{
	cache := mi.Find("Cached")
	return cache
}

func (mi Meminfo)MemUsed()*Storage{
	return mi.MemTotle().Sub(mi.MemFree()).Sub(mi.BufferCache())
}

func (mi Meminfo)BufferCache()*Storage{
	return mi.Buffer().Add(mi.Cache()).Add(mi.SReclaimable())
}

func (mi Meminfo)MemAvailable()*Storage{
	available := mi.Find("MemAvailable")
	return available
}

func (mi Meminfo)SwapTotle()*Storage{
	swapTotle := mi.Find("SwapTotal")
	return swapTotle
}

func (mi Meminfo)SwapFree()*Storage{
	sfree := mi.Find("SwapFree")
	return sfree
}

func (mi Meminfo)SwapUsed()*Storage{
	used := mi.SwapTotle().Sub(mi.SwapFree()).Sub(mi.SwapCache())
	return used
}

func (mi Meminfo)SwapCache()*Storage{
	scached := mi.Find("SwapCached")
	return scached
}

func (mi Meminfo)SReclaimable()*Storage{
	aslab := mi.Find("SReclaimable")
	return aslab
}


func NewMeminfo()Meminfo{
	info := make(map[string]*Storage)
	ami := string(readAll("/proc/meminfo"))
	lami := strings.Split(ami,"\n")
	for _,mi := range lami[:len(lami)-1]{
		spaceRe, _ := regexp.Compile(`:\s+`)
		lmi := spaceRe.Split(mi, -1)
		lcount := strings.Split(lmi[1]," ")


		count, err := strconv.ParseInt(lcount[0], 10, 64)
		if err !=nil{
			panic(err)
		}
		if len(lcount) == 2{
			switch lcount[1] {
			case "kB":
				count = count * KB
			}
		}
		info[lmi[0]] =  &Storage{ Size: count }
	}
	return Meminfo(info)
}


func (mi Meminfo)MIBMem()string{
	return fmt.Sprintf("MiB Mem :%9s total,%9s free,%9s used,%9s buff/cache",
		mi.MemTotle().MB(),mi.MemFree().MB(),mi.MemUsed().MB(),mi.BufferCache().MB())
}

func (mi Meminfo)MIBSwap()string{
	return fmt.Sprintf("MiB Swap:%9s total,%9s free,%9s used.%9s avail Mem",
		mi.SwapTotle().MB(),mi.SwapFree().MB(),mi.SwapUsed().MB(),mi.MemAvailable().MB())
}

func memSender(ch chan<- *Cell){
        y := 3
        if cpuDetail {
		onlineCpus := onlineCPU()
		y += len(onlineCpus) - 1
        }

	meminfo := NewMeminfo()
        cell := newCell()
        cell.msg = meminfo.MIBMem()
        cell.y = y
        ch <- cell

	cell = newCell()
	cell.msg = meminfo.MIBSwap()
	cell.y = y + 1
	ch <- cell
}
