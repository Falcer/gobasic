package main

import "fmt"

func main() {
	var array [3]int = [3]int{2, 12, 3}
	var slice []int
	fmt.Println(array[1])
	fmt.Println(len(array))

	//slice[4] = 5
	//fmt.Println(array[1])
	//fmt.Println(len(array))

	slice = append(slice, 3)
	fmt.Println(slice)
	fmt.Println(len(slice))

	slice = append(slice, 12)
	fmt.Println(slice)
	fmt.Println(len(slice))
}
