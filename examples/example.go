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
		var key int64
		now := time.Now().Unix()
		for i := 20; i > 0; i = i - 5 {
			for j := 0; j < 10; j++ {
				key = now + int64(i)
				wheel.AddTask(fmt.Sprintf("i: %d,j: %d", i, j), key)
				glog.V(4).Infof("input task %d\n", key)
			}
		}
	}()
	wheel.Run(func(tasks []string) {
		fmt.Println(tasks)
	})
}
