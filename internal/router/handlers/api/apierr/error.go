package apierr

// New Create An json compatable response error
func New(data interface{}) map[string]interface{} {
	return map[string]interface{}{"error": data}
}
