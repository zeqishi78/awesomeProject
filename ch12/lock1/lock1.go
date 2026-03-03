package main

import "sync"
import "sync/atomic"

var total int32
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, 1)
	}
}

func main() {

}
