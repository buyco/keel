package helper

import "maps"

// CopyMap copies a map in a new one
func CopyMap(originalMap map[string]any) map[string]any {
	newMap := make(map[string]any)
	maps.Copy(newMap, originalMap)
	return newMap
}
