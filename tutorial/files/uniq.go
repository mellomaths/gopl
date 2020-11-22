// Uniq is inspired by the Linux command line uniq that search for duplicated lines
package main

import (
	"bufio"
	"fmt"
	"os"
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

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	// Note: ignores potential errors of input.Err()
	for input.Scan() {
		line := input.Text()
		counts[line]++
	}
}

func printDuplicatedLines(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println("No file provided!")
	}

	for _, filename := range files {
		f := openFile(filename)
		if f == nil {
			continue
		}

		countLines(f, counts)
		f.Close()
	}

	printDuplicatedLines(counts)
}
