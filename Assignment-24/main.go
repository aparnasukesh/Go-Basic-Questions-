package main

import "fmt"

func main() {
	var string1, string2 string
	fmt.Printf("Enter your first string :")
	fmt.Scan(&string1)
	fmt.Printf("Enter your second string :")
	fmt.Scan(&string2)
	result := stringCompare(string1, string2)
	if result == -1 {
		fmt.Printf("%s is less than %s ", string1, string2)
	} else if result == 1 {
		fmt.Printf("%s is greater than %s", string1, string2)
	} else {
		fmt.Printf("%s is equal to %s", string1, string2)
	}
}
func stringCompare(str1, str2 string) int {
	if str1 < str2 {
		return -1
	} else if str1 > str2 {
		return 1
	}
	return 0
}
