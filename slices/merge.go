package main

import "fmt"

func main() {
	// slice := []int{00, 11, 22}
	// slice2 := []int{33, 44, 55}
	// slice2 = append(slice, slice2...)
	// fmt.Println(slice, "\n", slice2)

	// slice := []string{"one", "two", "three"}
	// slice2 := []string{"four", "five", "six"}
	// slice = append(slice, slice2...)
	// fmt.Println("This is the merge of strings -", slice)

	slice := []float64{1.3, 1.4, 1.5}
	slice2 := []float64{2.1, 2.2, 2.3}
	slice = append(slice2, slice...)
	fmt.Printf("this is the index 1- %0.2f, 2- %0.2f, 3-%0.2f, 4- %0.2f, 5- %0.2f, 6%0.2f from the first slice", slice[0], slice[1], slice[2], slice[3], slice[4], slice[5])

}
