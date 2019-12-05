package utils

// CopyMap copies a map in a new one
func CopyMap(originalMap map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, values := range originalMap {
		newMap[key] = values
	}
	return newMap
}
