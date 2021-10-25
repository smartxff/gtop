package main

import "fmt"

func top()string{
	head :=  fmt.Sprintf("top - %s up%s,%s,%s\n",timeNow(),uptime(),userCount(),loadavg()) 


	return head
}
