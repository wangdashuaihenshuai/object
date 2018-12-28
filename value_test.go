package object

import (
	"testing"
)

var data = `{
	"s1": "123123",
	"n1": 23123123,
	"b1": true,
	"m1": {
		"s2": "kkkkk",
		"n2": 324.123123,
		"b2": false
	},
	"a1": [
		"1212",
		{
			"s3": "zzzz",
			"n3": 12312.1231,
			"b3": true
		}
	]
}`

func TestGetMapString(t *testing.T) {
	object, err := NewValueByString(data)
	if err != nil {
		t.Error("init Value error")
		return
	}
	str, err := object.Key("s1").String()
	if err != nil {
		t.Error(err.Error())
	}

	if str != "123123" {
		t.Error("get map string key s1 value error")
	}
	_, err = object.Key("s1").Number()
	if err == nil {
		t.Error("should be type error")
	}

	_, err = object.Key("s1").Bool()
	if err == nil {
		t.Error("should be type error")
	}
	_, err = object.Key("s1").Map()
	if err == nil {
		t.Error("should be type error")
	}

	_, err = object.Key("s1").Index(2).Any()
	if err == nil {
		t.Error("should be type error")
	}

	_, err = object.Key("s1").Array()
	if err == nil {
		t.Error("should be type error")
	}

}

func TestGetMapNumber(t *testing.T) {
	object, err := NewValueByString(data)
	if err != nil {
		t.Error("init Value error")
	}
	num, err := object.Key("n1").Number()
	if err != nil {
		t.Error(err.Error())
	}
	if num != 23123123 {
		t.Error("get map number key n1 value error")
	}
}

func TestGetMapBool(t *testing.T) {
	object, err := NewValueByString(data)
	if err != nil {
		t.Error("init Value error")
	}
	b, err := object.Key("b1").Bool()
	if err != nil {
		t.Error(err.Error())
	}
	if !b {
		t.Error("get map string key s1 value error")
	}
}

func TestGetMap(t *testing.T) {
	object, err := NewValueByString(data)
	if err != nil {
		t.Error("init Value error")
	}
	_, err = object.Key("m1").Map()
	if err != nil {
		t.Error(err.Error())
	}
	_, err = object.Key("m1").Key("n2").Number()
	if err != nil {
		t.Error(err.Error())
	}
	_, err = object.Key("kk").Map()
	if err == nil {
		t.Error("get map not exist key should error")
	}
}

func TestGetArray(t *testing.T) {
	object, err := NewValueByString(data)
	if err != nil {
		t.Error("init Value error")
	}
	_, err = object.Key("a1").Array()
	if err != nil {
		t.Error(err.Error())
	}
	_, err = object.Key("a1").Index(0).String()
	if err != nil {
		t.Error(err.Error())
	}
	_, err = object.Key("a1").Index(2).String()
	if err == nil {
		t.Error("array get index should out of range")
	}
}
