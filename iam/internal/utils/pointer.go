package utils

// Pointer gets a pointer for any type.
func Pointer[T any](element T) *T {
	return &element
}
