package rushgo

import (
	"bytes"
	"encoding/json"
	"strings"
)

func ParseJSON(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ResponseBodyContains(responseBody []byte, searchStr string) bool {
	if bytes.Contains(bytes.ToLower(responseBody), []byte(strings.ToLower(searchStr))) {
		return true
	}
	return false
}
