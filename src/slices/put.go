package main

import "fmt"

func main() {
	// put := []string{"A", "B", "D", "E", "F"}
	// put = append(put, "")
	// fmt.Println(len(put), cap(put))
	// copy(put[3:], put[2:])
	// put[2] = "C"
	// fmt.Println(put)

	put:=[]int{1,2,4,5,6,7}
	fmt.Println(len(put),cap(put))
	copy(put[3:],put[2:])
	put[2]=3
	put=append(put,8)
	fmt.Println(put)

}
