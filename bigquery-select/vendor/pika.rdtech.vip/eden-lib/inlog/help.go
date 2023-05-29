package inlog

import "strconv"

func ConvertInt64ToString(m map[string]interface{}) (result map[string]interface{}) {
	for key, value := range m {
		switch v := value.(type) {
		case int64:
			int64s := strconv.FormatInt(v, 10)
			m[key] = int64s
		default:
			continue
		}

	}
	return m
}
