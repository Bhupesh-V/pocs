package main

import (
	"fmt"
)

func applyFunctionToElements[T any](fn func(T) T) func([]T) []T {
	return func(nums []T) []T {
		result := make([]T, len(nums))
		for i, v := range nums {
			result[i] = fn(v)
		}
		return result
	}
}

func doubleNumber(x int) int {
	return x * 2
}

func squareNumber(x int) int {
	return x * x
}

func tripleNumber(x int) int {
	return x * 3
}

func doubleString(x string) string {
	return x + x
}

func tripleString(x string) string {
	return x + x + x
}

func compose[T any](fns ...func(T) T) func([]T) []T {
	return func(nums []T) []T {
		result := nums
		for _, fn := range fns {
			result = applyFunctionToElements(fn)(result)
		}
		return result
	}
}

func main() {
	doubleAndSquareInt := compose(
		doubleNumber,
		squareNumber,
		tripleNumber,
	)
	nums := []int{1, 2, 3, 4, 5}
	resultInt := doubleAndSquareInt(nums)
	fmt.Println(resultInt) // Output: [12 48 108 192 300]

	doubleAndSquareStr := compose(doubleString, tripleString)
	numsStr := []string{"1", "2", "3", "4", "5"}
	resultStr := doubleAndSquareStr(numsStr)
	fmt.Println(resultStr) // Output: [111111 222222 333333 444444 555555]
}
