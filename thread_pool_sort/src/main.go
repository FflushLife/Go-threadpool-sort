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
	psort.Sort([]int{3, 4, 1, 8, 7})
	go pool.Create(5)
	time.Sleep(time.Second * 3)
	pool.SetCycle(false)
	time.Sleep(time.Second)
}
