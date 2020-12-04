package main

import (
	"errors"
	"fmt"
	"os"
)

func getStateNameByUf(uf string) (string, error) {
	type state struct {
		name string
		uf   string
	}
	stateList := [7]state{state{"Rio", "RJ"}, state{"São Paulo", "SP"}, state{"Amazonas", "AM"}, state{"Bahia", "BA"}}
	for i := 0; i < 7; i++ {
		st := stateList[i]
		if st.uf == uf {
			return st.name, nil // if state was found, error is nil
		}
	}

	return "", errors.New("state not found") // creates a new error in case of not finding the state name
}

func main() {
	for _, uf := range os.Args[1:] {
		stateName, err := getStateNameByUf(uf)
		if err != nil {
			fmt.Println("Error on getting state name of '" + uf + "': " + err.Error()) // logs the error
			continue

			// note that in this case the application will not stop
			// but in some cases you should use log.Fatal to stop the program with the error that was generated
		}

		fmt.Println(uf + " is " + stateName)
	}
}
