// Echo prints out everything you pass to it as the linux command line called echo
package main

import (
	"fmt"
	"os"
)

func echo(args []string) {
	var s, sep string // will be initialized as an empty string ("")
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}

func main() {
	echo(os.Args)
}
