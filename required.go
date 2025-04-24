package required

import (
	"fmt"
	"reflect"
)

func requiredValue(this reflect.Value) error {
	if this.Kind() == reflect.Pointer {
		this = this.Elem()
	}

	if this.Kind() == reflect.Struct {
		for i := 0; i < this.NumField(); i++ {
			field := this.Type().Field(i)
			_, ok := field.Tag.Lookup("required")
			isStruct := field.Type.Kind() == reflect.Struct || (field.Type.Kind() == reflect.Pointer && field.Type.Elem().Kind() == reflect.Struct)

			if ok {
				valField := this.Field(i)

				if valField.IsZero() {
					return fmt.Errorf("%q is required and not set", field.Name)
				}

				if isStruct {
					return requiredValue(valField)
				}
			}
		}
	}

	return nil
}

// Required takes a struct and identifies any struct members with the required
// struct tag (the value can be anything, but must exist for it to be a valid
// struct tag). If the value is the zero value, it will yield an error.
// Otherwise it will return nil.
func Required(this any) error {
	return requiredValue(reflect.ValueOf(this))
}
