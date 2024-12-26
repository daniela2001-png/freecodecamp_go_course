package hackerrankexercises

import (
	"fmt"
	"sort"
)

func SortByLen(letters []string) {
	sort.Slice(letters, func(i, j int) bool {
		if len(letters[i])%2 == 0 && len(letters[j])%2 == 0 { // both letters has an even len

			if len(letters[i]) == len(letters[j]) { // both letters has an equal len
				return letters[i] < letters[j] // should alphabetical order
			}
			return len(letters[i]) > len(letters[j]) // even letters should be after odd letters
		}
		if len(letters[i])%2 != 0 && len(letters[j])%2 != 0 { // both letters has an odd len
			if len(letters[i]) == len(letters[j]) { // both letters has an equal len
				return letters[i] < letters[j] // should alphabetical order
			}
			return len(letters[i]) < len(letters[j]) // odd letters must precede even letters
		}
		if len(letters[i])%2 != 0 { // odd letter must be first than a even letter
			return true // put it first
		} else {
			return false // else put it as second
		}
	})
}

func Solution(letters []string) []string {
	fmt.Println("letters: %v", letters)
	SortByLen(letters)
	fmt.Println("NEW custom string letters sort: %v", letters)
	return letters
}

// Solution([abc ab abcde a abcd abbd])  --> [a abc abcde abbd abcd ab]
