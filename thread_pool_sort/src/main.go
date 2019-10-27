package main

import (
	"fmt"
	"pool"
	"psort"
	"time"
)

func main() {
	fmt.Println("Start")
	go pool.Create(5)
	time.Sleep(time.Second * 3)
	pool.Cycle = false
	psort.Sort([]int{3, 4, 1})
	time.Sleep(time.Second)
}
