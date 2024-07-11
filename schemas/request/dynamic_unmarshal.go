package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
)

// Dyanmic UnmarshalJSON is created to ensure that JSON is correctly decoded into a struct.
// Also preventing the incoming request from HTTP is being bypassed because no error handler is available
func DynamicUnmarshalFromReader(reader io.Reader, target interface{}) error {
	var raw map[string]interface{}
	if err := json.NewDecoder(reader).Decode(&raw); err != nil {
		return err
	}

	targetValue := reflect.ValueOf(target).Elem()
	targetType := targetValue.Type()

	for i := 0; i < targetValue.NumField(); i++ {
		field := targetValue.Field(i)
		fieldType := targetType.Field(i)

		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		value, ok := raw[jsonTag]
		if !ok || isValueEmpty(value) {
			return fmt.Errorf("missing or empty field '%s' in JSON", jsonTag)
		}

		if err := setFieldValue(field, value); err != nil {
			return fmt.Errorf("error setting value for field '%s': %v", jsonTag, err)
		}
	}

	return nil
}

func isValueEmpty(value interface{}) bool {
	switch v := value.(type) {
	case nil:
		return true
	case string:
		return v == ""
	case float64:
		return false
	case bool:
		return false
	case map[string]interface{}:
		return len(v) == 0
	default:
		return false
	}
}

func setFieldValue(field reflect.Value, value interface{}) error {
	switch field.Kind() {
	case reflect.String:
		str, ok := value.(string)
		if !ok || str == "" {
			return errors.New("value is not a non-empty string")
		}
		field.SetString(str)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num, ok := value.(float64)
		if !ok {
			return errors.New("value is not a number")
		}
		field.SetInt(int64(num))

	case reflect.Float32, reflect.Float64:
		num, ok := value.(float64)
		if !ok {
			return errors.New("value is not a number")
		}
		field.SetFloat(num)

	case reflect.Bool:
		boolVal, ok := value.(bool)
		if !ok {
			return errors.New("value is not a boolean")
		}
		field.SetBool(boolVal)

	case reflect.Struct:
		// Recursive call for nested structs
		structVal, ok := value.(map[string]interface{})
		if !ok {
			return errors.New("value is not a struct")
		}
		if err := dynamicUnmarshalStruct(structVal, field); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unsupported field type: %v", field.Kind())
	}

	return nil
}

func dynamicUnmarshalStruct(raw map[string]interface{}, target reflect.Value) error {
	targetType := target.Type()

	for i := 0; i < target.NumField(); i++ {
		field := target.Field(i)
		fieldType := targetType.Field(i)

		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" {
			continue
		}

		value, ok := raw[jsonTag]
		if !ok || isValueEmpty(value) {
			return fmt.Errorf("missing or empty field '%s' in JSON", jsonTag)
		}

		if err := setFieldValue(field, value); err != nil {
			return fmt.Errorf("error setting value for field '%s': %v", jsonTag, err)
		}
	}

	return nil
}
