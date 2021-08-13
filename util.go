package toolbox

import "encoding/json"

func StructToMap(b interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m) // Convert to a map
	return m, err
}
