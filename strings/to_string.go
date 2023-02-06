package strings

func CastToString[T ~string](value T) string {
	return string(value)
}

func GetArrayOfStrings[T ~string](arr []T) []string {
	tmp := make([]string, len(arr))

	for i, val := range arr {
		tmp[i] = CastToString(val)
	}

	return tmp
}
