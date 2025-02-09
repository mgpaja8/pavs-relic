package maps

func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))

	for _, value := range m {
		values = append(values, value)
	}

	return values
}
