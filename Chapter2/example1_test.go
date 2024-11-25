package main

import (
	"fmt"
	"testing"
	"time"
)

func doWork(id int) {
	fmt.Printf("Work %d started at %s\n", id, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("Work %d finished at %s\n", id, time.Now().Format("15:04:05"))
}

func TestDoWork(t *testing.T) {
	for i := 0; i < 5; i++ {
		go doWork(i)
	}
	fmt.Println("Before sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("After sleep")
}
