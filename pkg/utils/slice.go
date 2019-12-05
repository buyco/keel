package utils

import (
	"github.com/thoas/go-funk"
	"reflect"
)

// InArray searches for a value from interface
func InArray(needle interface{}, haystack interface{}) (bool, int) {
	var exists = false
	var index = -1

	switch reflect.TypeOf(haystack).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(haystack)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return exists, index
			}
		}
	}
	return exists, index
}

// StringInSlice searches value in slice
func StringInSlice(searchVal string, list []string) bool {
	for _, elt := range list {
		if elt == searchVal {
			return true
		}
	}
	return false
}

// SliceDiff returns a diff between two slice
func SliceDiff(refKeys, toDelete []string) []string {
	toDelete = funk.Uniq(toDelete).([]string)
	tmpCompare := make(map[string]struct{}, len(toDelete))
	for _, x := range toDelete {
		tmpCompare[x] = struct{}{}
	}
	var diff []string
	for _, key := range refKeys {
		if _, found := tmpCompare[key]; !found {
			diff = append(diff, key)
		}
	}
	return diff
}
