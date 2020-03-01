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

func createTasks() [][]int {
	data_in := make([]byte, 8192*64)
	var tasks [][]int

	file, err := os.Open("rand_data.txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	count_bytes, err := file.Read(data_in)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	data := strings.Split(string(data_in[:count_bytes-1]), "\n")

	for _, el := range data {
		var task_single []int
		numbers_str := strings.Split(el, ",")
		for _, number_str := range numbers_str {
			number_64, err := strconv.Atoi(number_str)
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

func writeTime(ms int64, tCount uint64) {
	file, err := os.OpenFile("times.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	file.WriteString(strconv.FormatInt(ms, 10)+","+strconv.FormatUint(tCount, 10)+"\n")
}

func writeResult(result []int, counter int) {
	file, err := os.Create("results/result"+strconv.Itoa(counter)+".txt")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	for _, number := range result {
		file.WriteString(strconv.Itoa(number)+",")
	}
	file.WriteString("\n")
}

func main() {
	var sortInstance *psort.PSort
	var poolInstance *pool.Pool
	var tCount uint64 = 1
	var tasks [][]int

	if (len(os.Args) > 1) {
		tCount, _ = strconv.ParseUint(os.Args[1], 10, 64)
	}
	runtime.GOMAXPROCS(int(tCount))

	fmt.Println("Start initializing...")
	tasks = createTasks()
	fmt.Println("Tasks created")
	fmt.Println("Processing...")

	for i, target := range tasks {
		sortInstance = psort.New(target, tCount)
		poolInstance = pool.New(tCount, psort.TSort, unsafe.Pointer(sortInstance))
		poolInstance.ChangeTask(unsafe.Pointer(sortInstance))

		start := time.Now()
		poolInstance.Start();
		writeTime(time.Now().Sub(start).Nanoseconds() / 1000, tCount)
		fmt.Println("time=", time.Now().Sub(start))

		result := make([]int, 0)
		var sliceThreadSize float64 = float64(len(sortInstance.GetTarget())) / float64(tCount)
		// Merging
		for i := uint64(0); i < tCount; i++ {
			l := int(float64(i) * sliceThreadSize)
			r := int(float64(i + 1) * sliceThreadSize)
			result = psort.Merge(result, sortInstance.GetTarget()[l:r])
		}
		fmt.Println("final time=", time.Now().Sub(start))
		writeResult(result, i)
	}
}
