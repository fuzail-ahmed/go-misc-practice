package main

// check if string is palindrome or not
func isPalindrome(s string) bool {
	if s == "" || len(s) == 0 {
		return true
	}

	if len(s) == 1 {
		return true
	}

	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}
