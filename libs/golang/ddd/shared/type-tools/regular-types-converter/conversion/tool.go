package regulartypetool

import (
	"fmt"
	"reflect"
	"time"
)

var (
	dateLayout = "2006-01-02 15:04:05"
)

// ConvertFromMapStringToEntity converts a map[string]interface{} to the specified entity type.
// The entityType parameter should be the reflect.TypeOf() value of the target entity.
func ConvertFromMapStringToEntity(entityType reflect.Type, data map[string]interface{}) (interface{}, error) {
	if entityType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("entityType must be a struct")
	}

	entity := reflect.New(entityType).Elem()

	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		fieldName := field.Tag.Get("bson")
		if err := setFieldValue(entity, field, data[fieldName]); err != nil {
			return nil, fmt.Errorf("field %s: %w", fieldName, err)
		}
	}

	// Return a pointer to the newly created entity
	return entity.Addr().Interface(), nil
}

// setFieldValue sets the value of a field in a struct based on the provided data.
func setFieldValue(entity reflect.Value, field reflect.StructField, value interface{}) error {
	if value == nil {
		return fmt.Errorf("missing value")
	}

	fieldValue := entity.FieldByName(field.Name)
	if !fieldValue.IsValid() {
		return fmt.Errorf("invalid field")
	}

	switch fieldValue.Kind() {
	case reflect.Struct:
		return setStructField(fieldValue, field, value)
	case reflect.Slice:
		return setSliceField(fieldValue, field, value)
	default:
		return setBasicField(fieldValue, field, value)
	}
}

// setStructField sets the value of a struct field.
func setStructField(fieldValue reflect.Value, field reflect.StructField, value interface{}) error {
	if field.Type == reflect.TypeOf(time.Time{}) {
		return setTimeField(fieldValue, value)
	}

	nestedEntity, err := ConvertFromMapStringToEntity(field.Type, value.(map[string]interface{}))
	if err != nil {
		return err
	}
	fieldValue.Set(reflect.ValueOf(nestedEntity).Elem())

	return nil
}

// setSliceField sets the value of a slice field.
func setSliceField(fieldValue reflect.Value, field reflect.StructField, value interface{}) error {
	valueSlice, ok := value.([]interface{})
	if !ok {
		// If the value is not a slice of interface{}, it might be a slice of specific types
		valueSlice = make([]interface{}, reflect.ValueOf(value).Len())
		for i := 0; i < reflect.ValueOf(value).Len(); i++ {
			valueSlice[i] = reflect.ValueOf(value).Index(i).Interface()
		}
	}

	sliceValue := reflect.MakeSlice(field.Type, len(valueSlice), len(valueSlice))
	for i, item := range valueSlice {
		itemValue := sliceValue.Index(i)
		itemField := field.Type.Elem()

		if itemField.Kind() == reflect.Struct {
			switch v := item.(type) {
			case map[string]interface{}:
				nestedEntity, err := ConvertFromMapStringToEntity(itemField, v)
				if err != nil {
					return err
				}
				itemValue.Set(reflect.ValueOf(nestedEntity).Elem())
			default:
				if reflect.TypeOf(v).ConvertibleTo(itemField) {
					itemValue.Set(reflect.ValueOf(v).Convert(itemField))
				} else {
					return fmt.Errorf("unexpected type for slice element: %T", v)
				}
			}
		} else {
			itemValue.Set(reflect.ValueOf(item))
		}
	}

	fieldValue.Set(sliceValue)
	return nil
}

// setBasicField sets the value of a basic field (non-struct, non-slice).
func setBasicField(fieldValue reflect.Value, field reflect.StructField, value interface{}) error {
	fieldVal := reflect.ValueOf(value)
	if fieldVal.Type().ConvertibleTo(field.Type) {
		fieldValue.Set(fieldVal.Convert(field.Type))
		return nil
	}

	return fmt.Errorf("cannot convert %T to %s", value, field.Type)
}

// setTimeField sets the value of a time.Time field.
func setTimeField(fieldValue reflect.Value, value interface{}) error {
	switch v := value.(type) {
	case string:
		parsedTime, err := time.Parse(dateLayout, v)
		if err != nil {
			return err
		}
		fieldValue.Set(reflect.ValueOf(parsedTime))
	case time.Time:
		fieldValue.Set(reflect.ValueOf(v))
	default:
		return fmt.Errorf("expected string or time.Time, got %T", value)
	}
	return nil
}

// ConvertFromArrayMapStringToEntities converts an array of map[string]interface{} to an array of the specified entity type.
func ConvertFromArrayMapStringToEntities(entityType reflect.Type, dataArray []map[string]interface{}) ([]interface{}, error) {
	var entities []interface{}

	for _, data := range dataArray {
		entity, err := ConvertFromMapStringToEntity(entityType, data)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

func ConvertFromEntityToMapString(entity interface{}) (map[string]interface{}, error) {
	entityValue := reflect.ValueOf(entity)
	if entityValue.Kind() == reflect.Ptr {
		entityValue = entityValue.Elem()
	}
	fmt.Printf("Entity Value: %+v\n", entityValue) // Debugging line
	if entityValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("entity must be a struct")
	}

	entityType := entityValue.Type()
	data := make(map[string]interface{})

	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		if !field.IsExported() {
			continue // Skip unexported fields
		}
		fieldName := field.Tag.Get("bson")
		if fieldName == "" {
			fieldName = field.Name // Fallback to the field name if bson tag is not present
		}
		fieldValue := entityValue.Field(i)

		switch fieldValue.Kind() {
		case reflect.Struct:
			nestedData, err := ConvertFromEntityToMapString(fieldValue.Interface())
			if err != nil {
				return nil, err
			}
			data[fieldName] = nestedData
		case reflect.Slice:
			sliceData := make([]interface{}, fieldValue.Len())
			for j := 0; j < fieldValue.Len(); j++ {
				item := fieldValue.Index(j)
				if item.Kind() == reflect.Struct {
					nestedData, err := ConvertFromEntityToMapString(item.Interface())
					if err != nil {
						return nil, err
					}
					sliceData[j] = nestedData
				} else {
					sliceData[j] = item.Interface()
				}
			}
			data[fieldName] = sliceData
		default:
			data[fieldName] = fieldValue.Interface()
		}
	}

	return data, nil
}
