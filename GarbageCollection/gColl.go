package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats(mem)
	/*
		for i := 0; i < 10; i++ {
			s := make([]byte, 100000000)
			if s == nil {
				fmt.Println("Operation failed!")
			}
			time.Sleep(5 * time.Second)
		}
		printStats(mem)
	*/
	array := [...]int{2, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Println("unsafe.Sizeof(array[0]:", unsafe.Sizeof(array[0]))
	fmt.Println("pointer:", *pointer)
	fmt.Print(*pointer, " ")
	memoryAddress := uintptr(unsafe.Pointer(pointer)) +
		unsafe.Sizeof(array[0])
	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) +
			unsafe.Sizeof(array[0])

		fmt.Println("memoryAddress:", memoryAddress)
	}

	fmt.Println("memoryAddress:", memoryAddress)
	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) +
		unsafe.Sizeof(array[0])
	fmt.Println()
	fmt.Print(*pointer, " ")
}
