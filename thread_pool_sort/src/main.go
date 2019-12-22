package main

import (
	"fmt"
	"log"
	"os"
	"pool"
	"psort"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func createTasks() [][]uint64{
	data_in := make([]byte, 8192*64)
	var tasks [][]uint64

	file, err := os.Open("rand_data.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	_, err = file.Read(data_in)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	data := strings.Split(string(data_in), "\n")

	for _, el := range data {
		var task_single []uint64
		numbers_str := strings.Split(el, ",")
		for _, number_str := range numbers_str {
			number_64, err := strconv.ParseUint(number_str, 10, 64)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}

			task_single = append(task_single, number_64)
		}
		tasks = append(tasks, task_single)
	}
	return tasks
}

func main() {
	fmt.Println(createTasks())
	var sortInstance *psort.PSort
	var poolInstance *pool.Pool
	var tCount uint64 = 1

	if (len(os.Args) > 1) {
		tCount, _ = strconv.ParseUint(os.Args[1], 10, 64)
	}
	runtime.GOMAXPROCS(int(tCount))

	fmt.Println("Start initializing...")
	target := []int{3, 4, 1, 11, 8, 7, 2, 6, 5, 0, 9}
	for i := 30000000; i > 12; i-- {
		target = append(target, i)
	}
	target_copy := target
	sortInstance = psort.New(target, tCount)
	poolInstance = pool.New(tCount, psort.TSort, unsafe.Pointer(sortInstance))

	start := time.Now()
	poolInstance.Start();
	fmt.Println("time=", time.Now().Sub(start))

	result := make([]int, 0)
	var sliceThreadSize float64 = float64(len(sortInstance.GetTarget())) / float64(tCount)
	for i := uint64(0); i < tCount; i++ {
		l := int(float64(i) * sliceThreadSize)
		r := int(float64(i + 1) * sliceThreadSize)
		result = psort.Merge(result, sortInstance.GetTarget()[l:r])
	}
	fmt.Println("final time=", time.Now().Sub(start))
	fmt.Println(result[:100])

	result = make([]int, 0)
	sortInstance = psort.New(target_copy, tCount)
	poolInstance.ChangeTask(unsafe.Pointer(sortInstance))
	fmt.Println("new task started")
	poolInstance.Start()
	sliceThreadSize = float64(len(sortInstance.GetTarget())) / float64(tCount)
	for i := uint64(0); i < tCount; i++ {
		l := int(float64(i) * sliceThreadSize)
		r := int(float64(i + 1) * sliceThreadSize)
		result = psort.Merge(result, sortInstance.GetTarget()[l:r])
	}
	fmt.Println("final time=", time.Now().Sub(start))
	fmt.Println(result[:100])
}
