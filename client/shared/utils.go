package shared

func ConvertNilString(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}

func ConvertNilStringPointer(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}
