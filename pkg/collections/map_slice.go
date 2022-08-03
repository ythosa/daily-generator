package collections

// MapSlice returns a slice containing the results of applying the given transform function to each
// element in the original collection.
func MapSlice[T any, R any](slice []T, transform func(T) R) []R {
	if slice == nil {
		return nil
	}

	result := make([]R, len(slice))

	for i, element := range slice {
		result[i] = transform(element)
	}

	return result
}
