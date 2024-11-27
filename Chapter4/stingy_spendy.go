package main

import (
	"fmt"
	"time"
	"sync"
)

func stingy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		*money += 10
		mutex.Unlock()
	}

	fmt.Println("Stingy is done")
}

func spendy(money *int, mutex *sync.Mutex) {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		*money -= 10
		mutex.Unlock()
	}
	fmt.Println("Spendy is done")
}

func main() {
	money := 100
	mutex := sync.Mutex{}
	go stingy(&money, &mutex)
	go spendy(&money, &mutex)
	time.Sleep(1 * time.Second)
	mutex.Lock()
	fmt.Println(money)
	mutex.Unlock()
}