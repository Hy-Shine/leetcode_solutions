package utils

func MinStringLength(a, b string) int {
	if len(a) > len(b) {
		return len(b)
	}
	return len(a)
}

func MaxStringLength(a, b string) int {
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}
