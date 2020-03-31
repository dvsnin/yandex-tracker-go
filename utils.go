package tracker

// Преобразование типа в стринг
func interfaceToString(s interface{}) string {
	switch s.(type) {
	case string:
		return s.(string)
	default:
		return ""
	}
}
