package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jonyhy96/timer/pkg/timewheel"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	wheel := timewheel.NewTimeWheel(3600)
	go func() {
		key := time.Now().Unix() + 2
		glog.V(4).Infof("input task %d\n", key)
		wheel.AddTask("123", key)
	}()
	wheel.Run(func(tasks []string) {
		fmt.Println(tasks)
	})
}
