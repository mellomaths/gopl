// Echo prints out everything you pass to it as the linux command line called echo
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func echo(args []string) {
	start := time.Now()
	var s, sep string                // will be initialized as an empty string ("")
	for i := 1; i < len(args); i++ { // started on 1 because the first element is the name of this module
		s += sep + args[i]
		sep = " "
	}

	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Printf("%fs echo()\n", secs)
}

func echoSlice(args []string) {
	start := time.Now()
	s, sep := "", ""
	// range returns a tuple (index, value of the index inside the array)
	for _, arg := range args[1:] { // underscore is used because we don't need the index for the logic
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Printf("%fs echoSlice()\n", secs)
}

func echoJoin(args []string) {
	start := time.Now()
	fmt.Println(strings.Join(args[1:], " "))
	secs := time.Since(start).Seconds()
	fmt.Printf("%fs echoJoin()\n", secs)
}

func echoNotFormatted(args []string) {
	start := time.Now()
	fmt.Println(args[1:])
	secs := time.Since(start).Seconds()
	fmt.Printf("%fs echoNotFormatted()\n", secs)
}

func echoIndexes(args []string) {
	start := time.Now()
	for index, arg := range args[1:] {
		fmt.Println(strconv.Itoa(index) + ". " + arg)
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%fs echoIndexes()\n", secs)
}

func main() {
	echo(os.Args)
	echoSlice(os.Args)
	echoJoin(os.Args)
	echoNotFormatted(os.Args)
	echoIndexes(os.Args)
}
