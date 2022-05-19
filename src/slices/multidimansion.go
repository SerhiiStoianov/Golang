package main

import "fmt"

func main(){
	// firstTable := []float64{2.30, 3.29, 2.23}
	// secondTable := []float64{1.21, 2.23, 3.23}
	// multiSlice := [][]float64{firstTable, secondTable}
	// fmt.Println(multiSlice)

	// multiSlice := [][]float64{[]float64{1.23, 2.34, 3.34}, []float64{1.11, 2.53, 3.95}}
	// fmt.Println(multiSlice)

	multiSlice := [][]int{}
	for i := 0; i < 10; i++ {
		multiSlice = append(multiSlice, []int{i})
	}
	fmt.Println(multiSlice)
}