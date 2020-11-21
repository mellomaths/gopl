// Echo prints out everything you pass to it as the linux command line called echo
package main

import (
	"fmt"
	"os"
)

func echo(args []string) {
	var s, sep string // will be initialized as an empty string ("")
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}

	fmt.Println(s)
}

func echoSlice(args []string) {
	s, sep := "", ""
	// range returns a tuple (index, value of the index inside the array)
	for _, arg := range args[1:] { // underscore is used because we don't need the index for the logic
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func main() {
	echo(os.Args)
	echoSlice(os.Args)
}
