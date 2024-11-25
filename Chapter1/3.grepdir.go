// Change the program you wrote in the second exercise so that instead of passing
// a list of text filenames, you pass a directory path. The program will look inside
// this directory and list the files. For each file, you can spawn a goroutine that will
// search for a string match (the same as before). Call the program grepdir.go.
// Here’s how you can execute this Go program:
// go run grepdir.go bubbles ../../commonfiles
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
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
	dir := os.Args[2]
	substr := os.Args[1]
	root := os.DirFS(dir)
	fmt.Print(root)
	files, err := fs.Glob(root, "*")

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range files {
		filePath := path.Join(dir, v)
		go grepFile(substr, filePath)
	}
	time.Sleep(1 * time.Second)
}
