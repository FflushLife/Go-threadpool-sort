package pool

import "fmt"
import "time"

type cb func([]int, int)
var cycle, end = true, false
var callback cb

func wait(n int) {
	for cycle {
		fmt.Println("Waiting...")
		callback([]int{1,2,3}, n)
		time.Sleep(time.Second)
	}
	end = true
}

func Create(n int) {
	for i := 0; i < n; i++ {
		go wait(i)
	}
	fmt.Println("Pools have been created")
	for !end {}
	fmt.Println("End")
}

func SetCallback(f cb) {
	callback = f
}

func SetCycle(c bool) {
	cycle = c
}
