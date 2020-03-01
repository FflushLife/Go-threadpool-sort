package psort

import "unsafe"

type PSort struct {
	target []int
	tCount uint64
}

func New(target []int, tCount uint64) *PSort {
	return &PSort{target, tCount}
}

func (ps *PSort) GetTarget() []int {
	return ps.target
}

// Merge for master thread, not per goroutine
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

// Per goroutine sort
func TSort(uInstance unsafe.Pointer, n uint64) {
	var psInstance PSort = *(*PSort)(uInstance)
	var sliceThreadSize float64 = float64(len(psInstance.target)) /
								float64(psInstance.tCount)
	var l, r int

	l = int(float64(n) * sliceThreadSize)
	r = int(float64(n + 1) * sliceThreadSize)

	// Insert sorted result in resulted array
	tmp := append(psInstance.target[:l], mergesort(psInstance.target[l:r])...)
	psInstance.target = append(tmp, psInstance.target[r:]...)
}

// Sort engine
func mergesort(a []int) []int {
	var num = len(a)

	if num == 1 {
		return a
	}

	middle := int(num / 2)
	var (
		left = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = a[i]
		} else {
			right[i-middle] = a[i]
		}
	}

	return merge(mergesort(left), mergesort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left) + len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
