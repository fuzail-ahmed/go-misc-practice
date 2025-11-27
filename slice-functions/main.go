package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	SliceOps()
}

// CreateSlices function does not return anything.
//
// Created just to demonstrate the slice creation methods
func CreateSlices() {
	s := []int{}
	s1 := make([]int, 2)
	s2 := make([]int, 0, 10)

	fmt.Println(s, s1, s2)
}

// SliceOps
func SliceOps() {
	s := []string{"fuzail", "ahmed", "asrar", "ahmed", "ansari"}

	// basic slice
	fmt.Println(s[1:3])

	// from start
	fmt.Println(s[:3])

	// from start to end
	fmt.Println(s[2:])

	// append
	s = append(s, "alama", "alamgiri")

	s2 := []string{"asdflj"}
	s = append(s, s2...)
	fmt.Println(s)

	news := make([]string, len(s))
	fmt.Println(len(news), cap(news))
	copy(news, s)

	fmt.Println(s, news)

	// delete there is no build in function for deleting an element from an array
}

func deleteIndex(s []int, index int) []int {
	// convert 1, true, fuzail to json
	b1, _ := json.Marshal(1)
	b2, _ := json.Marshal(true)
	b3, _ := json.Marshal("fuzail")
	fmt.Println(string(b1), string(b2), string(b3))

	return []int{}
}
