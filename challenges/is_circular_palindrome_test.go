package challenges

import (
	"testing"
)

func throwErr(t *testing.T, expect interface{}, res interface{}) {
	t.Errorf("Expected: %v, Result: %v", expect, res)
}

func TestIsPalindromeSuccess(t *testing.T) {
	println("Is Palindrome Success Test Case")
	str := "ababa"
	res := isPalindrome(str)
	expect := true

	if res != expect {
		throwErr(t, expect, res)
	}
}

func TestIsNotPalindromeSuccess(t *testing.T) {
	println("Is Not Palindrome Success Test Case")
	str := "abcd"
	res := isPalindrome(str)
	expect := false

	if res != expect {
		throwErr(t, expect, res)
	}
}

func TestIsCircularPalindromeSuccessCases(t *testing.T) {
	println("Is Circular Palindrome Success Test Cases")
	strs := []string{"racecar", "mAlAyAlaM", "level", "rroto", "aaaad"}
	expect := true
	for _, str := range strs {
		res := IsCircularPalindrome(str)
		if res != expect {
			throwErr(t, expect, res)
		}
	}
}

func TestIsNotCircularPalindromeSuccessCases(t *testing.T) {
	println("Is Not Circular Palindrome Success Test Cases")
	strs := []string{"abc", "abcd", "hello world"}
	expect := false
	for _, str := range strs {
		res := IsCircularPalindrome(str)
		if res != expect {
			throwErr(t, expect, res)
		}
	}
}
