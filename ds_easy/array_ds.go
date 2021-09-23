package main

import (
	"fmt"
)

/*
 * Complete the 'reverseArray' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func reverseArray(a []int32) []int32 {
	l := len(a) - 1
	for f := 0; f < l; f++ {
		a[f] = a[f] - a[l]
		a[l] = a[f] + a[l]
		a[f] = a[l] - a[f]
		l -= 1
	}
	return a
}

func main() {
	arr := []int32{305, 97, 1290, 5591, 5930, 9317, 440, 6533, 7470}
	fmt.Println(reverseArray(arr))
}
