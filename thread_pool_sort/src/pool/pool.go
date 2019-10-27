package pool

import "fmt"
import "time"

var Cycle = true
var end = false

func wait() {
	for Cycle {
		fmt.Println("Waiting...")
		time.Sleep(time.Second)
	}
	end = true
}

func Create(n int) {
	fmt.Println("Pool created")
	for i := 0; i < n; i++ {
		go wait()
	}
	for !end {}
	fmt.Println("End")
}
