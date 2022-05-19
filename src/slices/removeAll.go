package main

import "fmt"

func main(){
	// removeAll:=[]int{1,2,3,4,5,6,7}
	// fmt.Println(removeAll)
	// removeAll=removeAll[:0]
	// fmt.Println(removeAll,len(removeAll),cap(removeAll))

	// r:=make([]int,6)
	// r[0]=1
	// r[1]=2
	// r[2]=3
	// r[3]=4
	// r[4]=5
	// r[5]=6
	// fmt.Println(r)
	// r=r[:0]
	// fmt.Println(r)
	// fmt.Println(len(r),cap(r))

	r:=[]int{1,2,3,4,5}
	r=nil
	fmt.Println(r)
	fmt.Println(len(r),cap(r))
}