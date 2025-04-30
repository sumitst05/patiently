package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// helper to create a snapshot of any struct as a JSON string
func SnapshotStruct(v any) (string, error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return "", fmt.Errorf("expected a pointer to a struct, got %s", val.Kind())
	}

	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
