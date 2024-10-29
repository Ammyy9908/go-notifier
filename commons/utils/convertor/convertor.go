package convertor

import (
	"encoding/json"
	"fmt"
)

func TypeConverter[R any](data any) (R, error) {
	var result R
	var b []byte
	var err error

	// Check if data is already a byte slice or string
	switch v := data.(type) {
	case []byte:
		b = v
	case string:
		b = []byte(v)
	default:
		// Marshal data if it's not already a JSON-compatible type
		b, err = json.Marshal(&data)
		if err != nil {
			fmt.Println("Error in Marshal", err)
			return result, err
		}
	}

	// Unmarshal into the result type
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println("Error in Unmarshal", err)
		return result, err
	}
	return result, nil
}
