package main

import "fmt"

//Map

func main() {

	//2.
	states := make(map[string]string)
	states["NY"] = "NewYork"
	states["CA"] = "California"
	states["TX"] = "Texas"
	fmt.Println(states)

	//City map Get data with the TX key name.
	Texas := states["TX"]
	fmt.Println("Selected city : ", Texas)

	//Delete data with CA key name
	delete(states, "CA")
	fmt.Println(states)

	states["NY"] = "NewYork"

	//Create a key list with the capacity of the States map object..
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	//Print the cities in states object according to index values in key list.
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
