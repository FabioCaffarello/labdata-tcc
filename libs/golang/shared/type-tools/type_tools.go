package typetools

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// ToString converts various types of data to a string
func ToString(data interface{}) (string, error) {
	switch v := data.(type) {
	case map[string]interface{}:
		return MapInterfaceToString(v), nil
	case map[string]string:
		return MapStringToString(v), nil
	case string:
		return v, nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v), nil
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case bool:
		return strconv.FormatBool(v), nil
	default:
		return "", fmt.Errorf("unsupported type: %T", v)
	}
}

// ToFloat64 converts various types of data to a float64
func ToFloat64(data interface{}) (float64, error) {
	switch v := data.(type) {
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	case int, int8, int16, int32, int64:
		return float64(v.(int)), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(v.(uint)), nil
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}
}

// MapInterfaceToString converts a map[string]interface{} to a sorted string
func MapInterfaceToString(m map[string]interface{}) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	str := ""
	for _, k := range keys {
		str += fmt.Sprintf("%s:%v;", k, m[k])
	}
	return str
}

// MapStringToString converts a map[string]string to a sorted string
func MapStringToString(m map[string]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	str := ""
	for _, k := range keys {
		str += fmt.Sprintf("%s:%s;", k, m[k])
	}
	return str
}

// ParseDateWithFormat parses a date string with a given format
func ParseDateWithFormat(date, format string) (time.Time, error) {
	parsedTime, err := time.Parse(format, date)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

// ParseBool parses a boolean string
func ParseBool(data string) (bool, error) {
	return strconv.ParseBool(data)
}
