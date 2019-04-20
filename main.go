package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	//go func() {
	//	fmt.Println("-----------")
	//}()

	go func() {
		//fmt.Println("-----")
		//for {
		//
		//}
	}()

	go func() {
		fmt.Println("-----------")
	}()

	//select {}

	//for {
	//
	//}

	time.Sleep(time.Second)
}
