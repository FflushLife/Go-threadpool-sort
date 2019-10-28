package psort

import "fmt"

var s []int

// Per thread sort
func TSort(n int) {
	fmt.Println("Callback %d", s[n])

}

func Sort(target []int) {
	fmt.Println("Sorting...")
	s = target
}
