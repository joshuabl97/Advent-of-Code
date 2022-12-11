package reusable

// returns value in alphabet if given a single digit string
// a = 1, b = 2, c = 3, A = 27, B = 28, C = 29, etc
func CharNum(s string) int {
	char := rune(s[0])
	if char >= 97 && char <= 122 {
		return int(char) - 96
	} else if char >= 65 && char <= 90 {
		return int(char) - 38
	}
	return 0
}

// checks if there are multiple of the same value in the same array of strings
// i.e ['a','a'] -> true ['a','b'] -> false
func DupeExist(arr []string) bool {
	visited := make(map[string]bool, 0)

	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] {
			return true
		} else {
			visited[arr[i]] = true
		}
	}

	return false
}

// iterates through a slice of a slice of strings and only returns only unique values on the inner slice
// i.e [['a','a','b','c','c'], ['a','b','b','c']] returns [['a','b','c'],['a','b','c']]
func NestedStringUnique(nestedS [][]string) [][]string {
	for i, s := range nestedS {
		nestedS[i] = UniqueStrSlice(s)
	}

	return nestedS
}

// reverses a slice of ints
// i.e [1,2,3] -> [3,2,1]
func ReverseIntSlice(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// inputs a slice of strings and returns a string slice containing only unique characters
// i.e ["a","b","b","b","c","c"] returns ["a","b","c"]
func UniqueStrSlice(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}
