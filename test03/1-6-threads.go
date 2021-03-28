package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime.NumCPU():", runtime.NumCPU())
	// 读取默认的线程数
	fmt.Println(runtime.GOMAXPROCS(0))
	// 设置线程数为 10
	runtime.GOMAXPROCS(100)
	// 读取当前的线程数
	fmt.Println(runtime.GOMAXPROCS(0))
}
