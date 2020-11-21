// Echo prints out everything you pass to it as the linux command line called echo
package main

import (
	"fmt"
	"os"
	"strings"
)

func echo(args []string) {
	var s, sep string                // will be initialized as an empty string ("")
	for i := 1; i < len(args); i++ { // started on 1 because the first element is the name of this module
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

func echoJoin(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}

func echoNotFormatted(args []string) {
	fmt.Println(args[1:])
}

func main() {
	echo(os.Args)
	echoSlice(os.Args)
	echoJoin(os.Args)
	echoNotFormatted(os.Args)
}
