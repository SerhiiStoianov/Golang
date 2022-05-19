package main

import "fmt"

func main() {
	// myMap := make(map[string]string)
	// myMap["one"] = "the first"
	// myMap["two"] = "the second"
	// myMap["three"] = "the third"
	// myMap["four"] = "the fourth"

	// fmt.Println(myMap)
	// fmt.Println(myMap["three"])
	// fmt.Println("The length is ", len(myMap))

	// myMap := make(map[int]string)
	// myMap[1] = "one"
	// myMap[2] = "two"
	// myMap[3] = "three"
	// myMap[4] = "four"

	// fmt.Println(myMap)

	// delete(myMap, 1)
	// fmt.Println(myMap)
	// fmt.Println("The length after deleting a value is ", len(myMap))

	// // когда нужно узнать, есть ли "key"

	// _, i := myMap[2]
	// fmt.Println(i)

	myShortMap := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	fmt.Println(myShortMap)
	fmt.Println("the lenght is", len(myShortMap))

}
