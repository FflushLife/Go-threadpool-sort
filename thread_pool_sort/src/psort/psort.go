package psort

import (
	"fmt"
	"math/rand"
	"unsafe"
)

type PSort struct {
	target []int
	tCount int
}

// Sort engine
func quicksort(a []int) []int {
    if len(a) < 2 {
        return a
    }

    left, right := 0, len(a)-1

    pivot := rand.Int() % len(a)

    a[pivot], a[right] = a[right], a[pivot]

    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }

    a[left], a[right] = a[right], a[left]

    quicksort(a[:left])
    quicksort(a[left+1:])

    return a
}

// Per thread sort
func TSort(uInstance unsafe.Pointer, n int) {
	fmt.Printf("Callback %d\n", n)

	var psInstance PSort = *(*PSort)(uInstance)
	var sliceThreadSize float64 = float64(len(psInstance.target)) / float64(psInstance.tCount)
	// TODO:: redo
	fmt.Println(quicksort(psInstance.target[int(float64(n) * sliceThreadSize):int(float64(n + 1) * sliceThreadSize)]))
}

func New(target []int, tCount int) PSort {
	fmt.Println("Sort inited")
	return PSort{target, tCount}
}
