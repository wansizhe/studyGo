package main

import (
	"fmt"
	"segmenttree"
)

func main() {
	arr := [...]int{1, 3, 5, 7, 9, 11}
	var st segmenttree.Segtree
	st.Build(arr[:], 6)
	st.PrintTree(15)
	// var size int = 6
	// build(arr[:], size)
	// // printTree(15)
	// fmt.Println(arr)
	// update(3, 8)
	// printTree(15)
	// fmt.Println(query(2, 4))
	// fmt.Println(arr)
	fmt.Printf("arr = %v\n", arr)
	s1 := arr[1:4]
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	s2 := append(s1, 100)
	s3 := append(s2, 200)
	fmt.Println(s1, s2, s3, arr)
	s4 := append(s3, 300)
	s4[0] = 400
	fmt.Println(s1, s2, s3, s4, arr)
}
