package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func printFile(filename string) {
    content, err := os.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(content))
}

func main() {
	fmt.Println(os.Args[1:])
	for _, i := range os.Args[1:] {
		go printFile(i)

	}
	time.Sleep(time.Second)
}

