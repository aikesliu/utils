package marshal

import "interface/json"

func Struct2JsonStr(v interface{}) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func JsonStr2Struct(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
