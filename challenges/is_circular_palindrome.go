package challenges

import (
	"strings"
)

// isPalindrome checks if the given string str is palindrome.
func isPalindrome(str string) bool {
	if len(str) < 2 {
		return false
	}
	rightIndex := len(str) - 1
	leftIndex := 0
	for leftIndex < rightIndex {
		if str[leftIndex] != str[rightIndex] {
			return false
		}
		leftIndex++
		rightIndex--
	}
	return true
}

// rotateString moves the first char in given string str to the back of string str.
func rotateString(str string) string {
	subStr1 := string(str[0])
	subStr2 := str[1:]
	return subStr2 + subStr1
}

// IsCircularPalindrome checks if the given string str is circular palindrome.
func IsCircularPalindrome(str string) bool {
	str = strings.ToLower(str)
	if isPalindrome(str) {
		return true
	}
	n := len(str)
	for i := 0; i < n-1; i++ {
		str = rotateString(str)
		if isPalindrome(str) {
			return true
		}
	}
	return false
}
