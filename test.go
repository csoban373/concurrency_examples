package main

import (
	"fmt"
	"time"
)

func main() {
	// numbers := make([]int, 1000000)
	// var waitGroup sync.WaitGroup
	// start := time.Now()
	// for i := range numbers {
	// 	waitGroup.Add(1)
	// 	go func(index int) {
	// 		time.Sleep(10000)
	// 		numbers[index] = index * 2
	// 		waitGroup.Done()
	// 	}(i)
	// }
	// waitGroup.Wait()
	// fmt.Println(time.Since(start).Nanoseconds())

	numbers := make([]int, 1000000)
	start := time.Now()
	for i := range numbers {
		time.Sleep(10000)
		numbers[i] = i * 2
	}
	fmt.Println(time.Since(start).Nanoseconds())
}

// numbers := make([]int, 1000000)
// 	start := time.Now()
// 	for i := range numbers {
// 		numbers[i] = i * 2
// 	}
// 	fmt.Println(time.Since(start).Nanoseconds())

// numbers := make([]int, 1000000)
// 	var waitGroup sync.WaitGroup
// 	start := time.Now()
// 	for i := range numbers {
// 		waitGroup.Add(1)
// 		go func(index int) {
// 			numbers[index] = index * 2
// 			waitGroup.Done()
// 		}(i)
// 	}
// 	waitGroup.Wait()
// 	fmt.Println(time.Since(start).Nanoseconds())

// numbers := make([]int, 0)
// 	for i := 0; i < 10; i++ {
// 		go func(index int) {
// 			numbers = append(numbers, index*2)
// 		}(i)
// 	}
// 	time.Sleep(100000)
// 	fmt.Println(numbers)

// nums := make([]int, 0)
// 	var mutex sync.Mutex
// 	for i := 0; i < 3; i++ {
// 		go func(index int) {
// 			mutex.Lock()
// 			nums = append(nums, index*2)
// 			mutex.Unlock()
// 		}(i)
// 	}
// 	time.Sleep(100000)
// 	fmt.Println(nums)

// numbers := make([]int, 10)
// 	var waitGroup sync.WaitGroup
// 	for i := range numbers {
// 		waitGroup.Add(1)
// 		go func(index int) {
// 			numbers[index] = index * 2
// 			waitGroup.Done()
// 		}(i)
// 	}
// 	waitGroup.Wait()
// 	fmt.Println(numbers)
