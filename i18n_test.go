package i18n

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewI18n(t *testing.T) {
	json_str := "{\"device\": \"1\",\"data\": [{\"humidity\": \"27\",\"time\": \"2017-07-03 15:23:12\"}]}"
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(json_str), &m)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m["device"])
		data := m["data"]
		if v, ok := data.([]interface{})[0].(map[string]interface{}); ok {
			fmt.Println(ok, v["humidity"], v["time"])
		}
	}
}
