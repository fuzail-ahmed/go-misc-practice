package main

// Write a program to remove duplicate element from array
func unique(arr []string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0, len(arr))

	for _, v := range arr {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}
