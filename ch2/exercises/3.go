/**
Write a program with three variables, one named b of type byte, one named smallI of type int32,
and one named bigI of type uint64. Assign each variable the maximum legal value for its type;
then add 1 to each variable. Print out their values
**/

package main

import "fmt"

func main() {
	var b byte = 255
	var smallI uint32 = 4_294_967_295
	var bigI uint64 = 18_446_744_073_709_551_615

	fmt.Println(b+1, smallI+1, bigI+1)
}
