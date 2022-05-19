package main

import "fmt"

func main(){
	// b:=[]int{2,3,4}
	// fmt.Println(b)
	// b=append([]int{1}, b...)
	// fmt.Println(b)

	s:=[]string{"two","three","four"}
	s=append([]string{"one"}, s...)
	fmt.Println(s)
}