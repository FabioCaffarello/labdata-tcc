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
		fieldName := field.Tag.Get("bson")
		if val, ok := data[fieldName]; ok {
			fieldValue := reflect.ValueOf(val)
			if field.Type.Kind() == reflect.Struct {
				nestedEntity, err := ConvertFromMapStringToEntity(field.Type, val.(map[string]interface{}))
				if err != nil {
					return nil, err
				}
				entity.Field(i).Set(reflect.ValueOf(nestedEntity))
			} else if fieldValue.Type().ConvertibleTo(field.Type) {
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

