package utils

import (
	"fmt"
	"reflect"
)

// helper for performing selective updates on update endpoints
// helps in updating only the fields that were sent in the payload
// prevents nullification of other fields

func UpdateStruct(original, updated any) error {
	originalVal := reflect.ValueOf(original).Elem()
	updatedVal := reflect.ValueOf(updated).Elem()

	// type checking
	if originalVal.Type() != updatedVal.Type() {
		return fmt.Errorf("mismatched types: %s and %s", originalVal.Type(), updatedVal.Type())
	}

	for i := range updatedVal.NumField() {
		// get  field names and values
		field := updatedVal.Type().Field(i)
		updatedField := updatedVal.Field(i)

		// update only non-zero fields
		if !updatedField.IsZero() {
			originalField := originalVal.FieldByName(field.Name)
			if originalField.IsValid() && originalField.CanSet() {
				originalField.Set(updatedField)
			}
		}
	}

	return nil
}
