package main

import (
	"fmt"
	"sort"
)

func nonempty(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a[:])
	//reverse(a[:])
	fmt.Println(a[:])
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))

	ages := make(map[string]int)
	ages["carol"] = 21 // panic: assignment to entry in nil map
	age, ok := ages["bob"]
	if !ok { /* "bob" is not a key in this map; age == 0. */
		fmt.Printf("print age %d and %t", age, ok)
	}
}
