package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slices []int) {
	for i, slice := range slices{
		slices[i] = f(slice)
	}
}

func mapArray(f func(a int) int, array [3]int) {
	for i, elem := range array {
		array[i] = f(elem)
	}

}

func main() {
	intsSlice := make([]int, 3)
	intsSlice[0] = 1
	intsSlice[1] = 2
	intsSlice[2] = 3

	mapSlice(addOne, intsSlice)
	fmt.Println(intsSlice)

	var intsArray [3]int
	intsArray[0] = 7
	intsArray[1] = 8
	intsArray[2] = 9

	mapArray(addOne, intsArray)
	fmt.Println(intsArray)

}
