package main

import "fmt"

func sum(numbers ...int) int{
	r := 0
	for _ , number := range numbers {
		r = r + number
	}
	return r
}

func main(){
	//var inhaArray [7] string //array
//	var inhaSlice [] string //slice
//	inhaSlice = make([]string, 7)
	inhaSlice := []int{1,2,3,4,5}

//	inhaSlice[0] = "i"
//	inhaSlice[1] = "n"
//	inhaSlice[2] = "h"
//	inhaSlice[3] = "a"

	for i := 0 ; i < len(inhaSlice); i++ {
		fmt.Println(inhaSlice[i])
	}
	fmt.Println(sum(inhaSlice...))
	

//	fmt.Println(1)
}
