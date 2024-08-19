package utils

func Filter[T any](input []T, fn func(T) bool) (result []T) {
	for _, v := range input {
		if fn(v) {
			result = append(result, v)
		}
	}

	return
}

func Contains[T comparable](input []T, value T) bool {
	for _, v := range input {
		if v == value {
			return true
		}
	}

	return false
}

func Remove[T comparable](input []T, value T) (result []T) {
	result = input

	for i, v := range input {
		if v == value {
			result = append(input[:i], input[i+1:]...)
		}
	}

	return
}
