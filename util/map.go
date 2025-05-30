package util 

func MapContains[K comparable, T any](m map[K]T, key K) bool {
	_, ok := m[key]
	return ok
}
