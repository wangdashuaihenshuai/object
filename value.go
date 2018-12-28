package object

import (
	"encoding/json"
	"fmt"
)

func NewValue(value interface{}) *Value {
	return &Value{value: value, prefix: "value", err: nil}
}

func NewValueByString(data string) (*Value, error) {
	var object interface{}
	err := json.Unmarshal([]byte(data), &object)
	if err != nil {
		return nil, err
	}
	return NewValue(object), nil
}

type Value struct {
	value  interface{}
	prefix string
	err    error
}

func (v *Value) Index(i int) *Value {
	if v.err != nil {
		return v
	}
	arr, ok := v.value.([]interface{})
	if !ok {
		v.err = fmt.Errorf("%s type is not []interface{}", v.prefix)
		return v
	}
	if i >= len(arr) {
		v.err = fmt.Errorf("%s[%d] index out of range", v.prefix, i)
		return v
	}
	newValue := arr[i]
	newPrefix := fmt.Sprintf("[%d]", i)
	return &Value{value: newValue, prefix: v.prefix + newPrefix, err: nil}
}

func (v *Value) Key(k string) *Value {
	if v.err != nil {
		return v
	}
	m, ok := v.value.(map[string]interface{})
	if !ok {
		v.err = fmt.Errorf("%s type is not map[string]interface{}", v.prefix)
		return v
	}
	newValue, ok := m[k]
	if !ok {
		v.err = fmt.Errorf("%s.%s is not exist", v.prefix, k)
		return v
	}
	newPrefix := fmt.Sprintf(".%s", k)
	return &Value{value: newValue, prefix: v.prefix + newPrefix, err: nil}
}

func (v *Value) Number() (float64, error) {
	if v.err != nil {
		return 0, v.err
	}
	switch i := v.value.(type) {
	case uint:
		return float64(i), nil
	case uint8:
		return float64(i), nil
	case uint16:
		return float64(i), nil
	case uint32:
		return float64(i), nil
	case uint64:
		return float64(i), nil
	case float32:
		return float64(i), nil
	case float64:
		return float64(i), nil
	}
	return 0, fmt.Errorf("%s type is not number", v.prefix)
}

func (v *Value) String() (string, error) {
	if v.err != nil {
		return "", v.err
	}

	str, ok := v.value.(string)
	if !ok {
		return "", fmt.Errorf("%s type is not string", v.prefix)
	}
	return str, nil
}

func (v *Value) Bool() (bool, error) {
	if v.err != nil {
		return false, v.err
	}

	b, ok := v.value.(bool)
	if !ok {
		return false, fmt.Errorf("%s type is not bool", v.prefix)
	}
	return b, nil
}

func (v *Value) Map() (map[string]interface{}, error) {
	if v.err != nil {
		return nil, v.err
	}
	m, ok := v.value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("%s type is not map[string]interface{}", v.prefix)
	}
	return m, nil
}

func (v *Value) Array() ([]interface{}, error) {
	if v.err != nil {
		return nil, v.err
	}
	arr, ok := v.value.([]interface{})
	if !ok {
		return nil, fmt.Errorf("%s type is not []interface{}", v.prefix)
	}
	return arr, nil
}

func (v *Value) Any() (interface{}, error) {
	if v.err != nil {
		return nil, v.err
	}
	return v.value, nil
}

func main() {
	const data = `
	{
  "meta": {
    "rooted": false,
    "os": "iOS",
    "bundle_version": "1.0",
    "rom": "iOS",
    "device_id": "12312",
    "bundle_id": "org.cocoapods.demo.APFAnswers-Example",
    "model": "x86_64",
    "os_version": "11.4"
  },
  "logs": [
    {
      "type": "launch",
      "log_id": 422,
      "timestamp": 1542357107415,
      "params": {
        "total": 1000,
        "pre_main": 100
      }
    },
    {
      "type": "page",
      "log_id": 423,
      "timestamp": 1542357107416,
      "params": {
        "page_name": "page_name0",
        "time": 300,
        "status": 1
      }
		}]
	}
	`

	var object interface{}
	err := json.Unmarshal([]byte(data), &object)
	if err != nil {
		panic(err)
	}
	value := NewValue(object)
	meta, err := value.Key("meta").Map()
	if err != nil {
		panic(err)
	}
	fmt.Println(meta["os"])
	logs, err := value.Key("logs").Array()
	if err != nil {
		panic(err)
	}
	fmt.Println(logs[0])
	log, err := value.Key("logs").Index(10).Map()
	if err != nil {
		panic(err)
	}
	fmt.Println(log["type"])
}
