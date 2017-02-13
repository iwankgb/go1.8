package main

import (
	"fmt"
	"time"
)

func main() {
	aMap := map[string]int{"one": 1, "two": 2, "three": 3}

	for v, k := range aMap {
		fmt.Sprintf("%s=>%s\n", v, k)
	}

	go printMe(aMap)
	go writeMe(&aMap)
	go printMe(aMap)
	go writeMe(&aMap)
	go printMe(aMap)
	go writeMe(&aMap)
	go printMe(aMap)
	go writeMe(&aMap)
	go printMe(aMap)
	go writeMe(&aMap)
	go printMe(aMap)
	go writeMe(&aMap)
	go printMe(aMap)
	go writeMe(&aMap)
	time.Sleep(1 * time.Second)
}

func printMe(aMap map[string]int) {
	for v, k := range aMap {
		fmt.Sprintf("%s=>%s\n", v, k)
	}
}

func writeMe(aMap *map[string]int) {
	realMap := *aMap
	realMap["random zero"] = 0
}
