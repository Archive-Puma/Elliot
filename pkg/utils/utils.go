package utils

// === PUBLIC METHODS ===

// FilterDuplicates removes duplicates from a list
func FilterDuplicates(data []string) []string {
	var unique []string
	duplicates := make(map[string]int)
	// Iterate over all the data
	for _, d := range data {
		duplicates[d]++
		if duplicates[d] == 1 {
			unique = append(unique, d)
		}
	}
	return unique
}
