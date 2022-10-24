package prioqueue

// Ascending order, or "min heap".
func Asc[P Number](a, b P) bool {
	return a < b
}

// Descending order, or "max heap".
func Desc[P Number](a, b P) bool {
	return a > b
}
