// Adapt the program in the third exercise to continue searching recursively in
// any subdirectories. If you give your search goroutine a file, it should search for a
// string match in that file, just like in the previous exercises. Otherwise, if you
// give it a directory, it should recursively spawn a new goroutine for each file or
// directory found inside. Call the program grepdirrec.go, and execute it by running this command:
// go run grepdirrec.go bubbles ../../commonfiles
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GrepFile(substring string, fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Println(err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, substring) {
			fmt.Printf("%s Content %s\n", strings.Repeat("-", 20), strings.Repeat("-", 20))
			fmt.Printf("%s contains %s\n\n", fileName, substring)
			break
		}
	}
}

func main() {
	substr := os.Args[1]
	dir := os.Args[2]
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			infor, err := os.Stat(path)
			if !infor.IsDir() {
				go GrepFile(substr, path)
			}
			return nil
		})

	if err != nil {
		log.Fatal(err)
	}
}