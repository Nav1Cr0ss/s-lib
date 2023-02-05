package enum

func GetArrayOfStrings[T ~string](A []T) []string {
	var tmp []string

	for _, val := range A {
		tmp = append(tmp, string(val))
	}

	return tmp
}
