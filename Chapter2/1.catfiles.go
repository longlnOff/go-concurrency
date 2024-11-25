// Write a program similar to the one in listing 2.3 that accepts a list of text file names as arguments.
// For each filename, the program should spawn a new
// goroutine that will output the contents of that file to the console. You can use
// the time.Sleep() function to wait for the child goroutines to complete (until
// you know how to do this better). Call the program catfiles.go. Here’s how you
// can execute this Go program:
// go run catfiles.go txtfile1 txtfile2 txtfile3

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func readFile(fileName string) []byte {
	fmt.Printf("%s Content %s\n", strings.Repeat("-", 20), strings.Repeat("-", 20))
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
	fmt.Printf("%s End %s\n", strings.Repeat("-", 20), strings.Repeat("-", 20))
	return data
}

func main() {
	argsWithoutProg := os.Args[1:]
	for _, arg := range argsWithoutProg {
		go readFile(arg)
	}
	time.Sleep(1 * time.Second)
}
