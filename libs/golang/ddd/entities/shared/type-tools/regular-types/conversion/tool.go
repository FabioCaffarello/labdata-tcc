package regulartypetool

import (
	"fmt"
	"reflect"
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
		fieldName := field.Tag.Get("json")
		if val, ok := data[fieldName]; ok {
			fieldValue := reflect.ValueOf(val)
			if fieldValue.Type().ConvertibleTo(field.Type) {
				entity.Field(i).Set(fieldValue.Convert(field.Type))
			} else {
				return nil, fmt.Errorf("field %s has invalid type", fieldName)
			}
		} else {
			return nil, fmt.Errorf("field %s is missing in the data", fieldName)
		}
	}

	return entity.Interface(), nil
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
