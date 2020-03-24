package p

func StrOrEmpty(input *string) string {
	if input == nil {
		return ""
	}
	return *input
}
