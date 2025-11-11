/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-11
 * @fileoverview This program calculates the sum of the numbers 10, -20, and 85.
 */

package main

import "fmt"

func main() {
	// Declaring variables
	const num1 int = 10;
	const num2 int = -20;
	const num3 int = 85;

	// Calculating the answer
	const answer int = num1 + num2 + num3;

	// Printing
	fmt.Println("10 + (-20) + 85 = " + fmt.Sprint(answer));
}