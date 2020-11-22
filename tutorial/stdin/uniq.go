// Uniq is inspired by the Linux command line uniq that search for duplicated lines
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// Note: ignores potential errors of input.Err()
	for input.Scan() {
		line := input.Text()
		if len(line) == 0 {
			break
		}

		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
