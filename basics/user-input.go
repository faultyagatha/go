package main

import "fmt"

/* Ex1: Write a program which prompts the user to enter a floating point number
and prints the integer which is a truncated version of the floating
point number that was entered. Truncation is the process of removing
the digits to the right of the decimal place. */

func main() {
	var appleNum int
	fmt.Printf("Number of apples?")

	fmt.Scan(&appleNum)
	fmt.Println("Apples number is: ", appleNum)
}

/* Ex2: Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter
if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings,
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!”
for the following strings, “ihhhhhn”, “ina”, “xian”. */
