package prioqueue

// Ascending order
func Asc[P Number](a, b P) bool {
	return a < b
}

// Descending order
func Desc[P Number](a, b P) bool {
	return a > b
}
