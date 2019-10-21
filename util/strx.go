package util

import "encoding/json"

func UnmarshalMap(str string) (map[string]string, error) {
	if str == "" {
		return make(map[string]string), nil
	}
	var m map[string]string
	e := json.Unmarshal([]byte(str), &m)
	if e != nil {
		return nil, e
	}
	return m, nil
}

func MarshalMap(m map[string]string) (string, error) {
	b, e := json.Marshal(m)
	if e != nil {
		return "", e
	}
	return string(b), nil
}
