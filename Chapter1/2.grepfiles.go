// Expand the program you wrote in the first exercise so that instead of printing
// the contents of the text files, it searches for a string match. The string to search
// for is the first argument on the command line. When you spawn a new goroutine,
// instead of printing the file’s contents, it should read the file and search for
// a match. If the goroutine finds a match, it should output a message saying that
// the filename contains a match. Call the program grepfiles.go. Here’s how you
// can execute this Go program (“bubbles” is the search string in this example):
// go run grepfiles.go bubbles txtfile1 txtfile2 txtfile3
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func grepFile(substring string, fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, substring) {
			fmt.Printf("%s Content %s\n", strings.Repeat("-", 20), strings.Repeat("-", 20))
			fmt.Printf("%s contains %s\n", fileName, substring)
			fmt.Printf("%s End %s\n", strings.Repeat("-", 20), strings.Repeat("-", 20))
			break
		}
	}
}

func main() {
	argsWithoutProg := os.Args[2:]
	substring := os.Args[1]
	fmt.Printf("Searching for `%s`\n", substring)
	for _, arg := range argsWithoutProg {
		go grepFile(substring, arg)
	}
	time.Sleep(4 * time.Second)
}
