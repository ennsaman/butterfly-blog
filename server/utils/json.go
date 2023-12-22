package utils

import "encoding/json"

const JSON_UTIL_ERR_PREFIX = "utils/json.go -> "

var Json = new(_json)

type _json struct{}

// Marshal data to json string
func (*_json) Marshal(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		Logger.Error(JSON_UTIL_ERR_PREFIX + "Marshal: " + err.Error())
		panic(err)
	}
	return string(data)
}

// Unmarshal json string to data
func (*_json) Unmarshal(data string, v any) {
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		Logger.Error(JSON_UTIL_ERR_PREFIX + "Unmarshal: " + err.Error())
		panic(err)
	}
}
