package main

import (
	"fmt"
	"pool"
	"psort"
	"time"
)

func main() {
	fmt.Println("Start")
	pool.SetCallback(psort.TSort)
	go pool.Create(5)
	time.Sleep(time.Second * 3)
	pool.SetCycle(false)
	psort.Sort([]int{3, 4, 1})
	time.Sleep(time.Second)
}
