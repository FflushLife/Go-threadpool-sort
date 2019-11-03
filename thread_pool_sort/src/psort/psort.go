package psort

import (
	"fmt"
	"math/rand"
	"unsafe"
)

type PSort struct {
	target []int
	tCount uint64
}

func New(target []int, tCount uint64) *PSort {
	fmt.Println("Sort inited")
	return &PSort{target, tCount}
}

func Merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

// Per thread sort
func TSort(uInstance unsafe.Pointer, n uint64) {
	fmt.Printf("Callback %d, len=%d\n", n, len((*(*PSort)(uInstance)).target))

	var psInstance PSort = *(*PSort)(uInstance)
	var sliceThreadSize float64 = float64(len(psInstance.target)) / float64(psInstance.tCount)
	var l, r int
	// TODO:: redo
	l = int(float64(n) * sliceThreadSize)
	r = int(float64(n + 1) * sliceThreadSize)
	fmt.Println(quicksort(psInstance.target[l:r]))
	tmp := append(psInstance.target[:l], quicksort(psInstance.target[l:r])...)
	psInstance.target = append(tmp, psInstance.target[r:]...)
	fmt.Println(psInstance.target)
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
