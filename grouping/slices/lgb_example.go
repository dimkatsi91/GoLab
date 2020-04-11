/* Little Go Book Example : Pg. 44
 *
 *	Efficient way to remove a value from an unsorted slice
 */

 package main

 import "fmt"


func removeAtIndex(src []int, idx int) []int {
	lastIndex := len(src) - 1
	// Swap the last value and the value we want to remove
	src[idx], src[lastIndex] = src[lastIndex], src[idx]
	// return from the start to lastIndex
	return src[:lastIndex]
}

 func main() {
	 scores := []int{1,2,3,4,5}						// [1 2 3 5 4]
	 scores = removeAtIndex(scores, 2)			   //  remove '3'
	 fmt.Println(scores)						  //   [1 2 5 4]
 }
 