package main

import (
	"fmt"
)

func function01() {
	var a string = "hello world"
	// var b = "Hi"
	// var c string
	// var d, e, f int
	// var g, h, i = true, 9.2, "Next" //bool, float64, string
	// j := 100
	// k, l := 1, 2

	// //
	// const ip string = "127.0.0.1"
	// const apple, banana = 12, 8
	// const (
	// 	Unknow = 0
	// 	Male   = 1
	// 	Female = 2
	// )

	//
	fmt.Println(len(a))
	fmt.Println(a[0], a[7])
}

// btoi true->1 and flase->0
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// itob whether i is non-zero
func itob(i int) bool { return i != 0 }

// Array, Slice
func test() {
	var q [3]int = [3]int{1, 2, 3} //Array
	var r [3]int = [3]int{1, 2}    //Slice
	fmt.Println(r[2])              //"0"
	var u := [...]int{1, 2, 3}
	fmt.Println("%T\n,q")
}
