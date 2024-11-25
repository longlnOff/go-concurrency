// Modify our sequential-letter frequency program to produce a list of word frequencies rather than letter frequencies. You can use the same URLs for the RFC
// web pages as were used in listing 3.3. Once it’s finished, the program should
// output a list of words with the frequency with which each word appears in the
// web page. Here’s some sample output:
// $ go run wordfrequency.go
// the -> 5
// a -> 8
// car -> 1
// program -> 3
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func countWords(url string, frequency map[string]int) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	words := strings.Fields(string(b))
	for _, word := range words {
		frequency[word]++
	}
}

func main() {
	var frequency = make(map[string]int)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		countWords(url, frequency)
		// If using goroutines, errors happen
		// go countWords(url, frequency)
	}

	time.Sleep(5 * time.Second)
	for word, count := range frequency {
		fmt.Printf("%s -> %d\n", word, count)
	}
}
