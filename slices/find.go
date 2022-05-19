package main

import "fmt"

func main() {
	// find := []string{"John", "Shin", "Gold", "Silver"}
	// for i, word := range find {
	// 	if word == "Gold" {
	// 		fmt.Println("Gold found at index", i)
	// 	}
	// }

	// iterat := make([]string, 3)
	// iterat[0] = "song"
	// iterat[1] = "bofy"
	// iterat[2] = "pind"
	// fmt.Println(iterat)
	// for i := 0; i < len(iterat); i++ {
	// 	fmt.Println(i, iterat[i])
	// }

	iterate := make([]int, 5)
	fmt.Println(iterate)
	for i := 0; i < len(iterate); i++ {
		if iterate[i] == 1 {
			fmt.Println("you're right")
		} else {
			fmt.Println("it was interesting")
		}
	}
}
