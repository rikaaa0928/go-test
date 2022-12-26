package go_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIntToString(t *testing.T) {
	str := `{"a":"b","b":1}`
	m := make(map[string]interface{})
	json.Unmarshal([]byte(str), &m)
	for k, v := range m {
		t.Logf("%s:%s\n", k, fmt.Sprintf("%v", v))
	}
	t.Logf("%v", m)
}
