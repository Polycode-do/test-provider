package shared

// `ConvertNilString` converts a string to a string pointer, returning nil if the string is empty.
// @param {string} str - The string to convert.
// @returns {string} - The converted string.
func ConvertNilString(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}

// `ConvertNilStringPointer` converts a string pointer to a string, returning an empty string if the pointer is nil.
// @param {string} str - The string to convert.
// @returns {string} - The converted string.
func ConvertNilStringPointer(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}
