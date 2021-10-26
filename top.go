package main

import "fmt"

func head()string{
	return fmt.Sprintf("top - %s up%s,%s,%s\n",timeNow(),uptime(),userCount(),loadavg())
}

func tasks()string{
	pt := NewProcTab()
        pt.Refresh()
	return fmt.Sprintf("Tasks: %d total,   %d running, %d sleeping,   %d stopped,   %d zombie",
		maxtask,running,sleepin,stopped,zombied)
}
