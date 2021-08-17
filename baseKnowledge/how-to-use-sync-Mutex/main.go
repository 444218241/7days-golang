package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex
var set = make(map[int]bool)

func main() {
	for i := 0; i < 100; i++ {
		go printOnce(100)
	}
	time.Sleep(2 * time.Second)
}

func printOnce(num int) {
	m.Lock()
	defer m.Unlock()
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
}
