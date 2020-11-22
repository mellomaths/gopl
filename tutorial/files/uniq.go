// Uniq is inspired by the Linux command line uniq that search for duplicated lines
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// OpenFile uses os.Open() to open a file
func openFile(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "uniq: %v\n", err)
		return nil
	}

	return f
}

// ReadFile uses ioutil.ReadFile() to read a file
func readFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "uniq: %v\n", err)
		return nil
	}

	return data
}

func printDuplicatedLines(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	// TODO: Store the name of the file that the line appeared and show a list of files that duplicated certain line

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("No file provided!")
	}

	fmt.Println("Executing with os.Open()")
	for _, filename := range files {
		f := openFile(filename)
		if f == nil {
			continue
		}

		input := bufio.NewScanner(f)
		// Note: ignores potential errors of input.Err()
		for input.Scan() {
			line := input.Text()
			counts[line]++
		}
		f.Close()
	}

	printDuplicatedLines(counts)

	counts = make(map[string]int) // Resets the map to count everything again

	fmt.Println("Executing with ioutil.ReadFile()")
	for _, filename := range files {
		data := readFile(filename)
		stringData := string(data)               // Transformes []byte to string
		lines := strings.Split(stringData, "\n") // Split the string on every line
		for _, line := range lines {
			if len(line) != 0 {
				counts[line]++
			}
		}
	}

	printDuplicatedLines(counts)
}
