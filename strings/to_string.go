package strings

func CastToString[T ~string](value T) string {
	return string(value)
}

func GetArrayOfStrings[T ~string](arr []T) []string {
	tmp := make([]string, len(arr))

	for _, val := range arr {
		tmp = append(tmp, CastToString(val))
	}

	return tmp
}
