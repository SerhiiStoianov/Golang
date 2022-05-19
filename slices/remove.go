package main

import "fmt"

func main() {
	// remove := []string{"a", "b", "c", "d", "f", "e", "g"}
	// cope := make([]string, len(remove))
	// copy(cope, remove)
	// fmt.Println(remove)
	// fmt.Println(cope)

	// cope = append(cope[:3], cope[4:]...)
	// fmt.Println(cope)

	remove := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(remove)
	remove = append(remove[:2], remove[3:]...)
	fmt.Println(remove)
}
