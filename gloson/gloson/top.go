package gloson

import "gloson/fun"

//反序列化到map
func Map(json string) (map[string]interface{}, error) {
	var um fun.Unmarshall
	return um.Unmarshall2Map(json)
}

//反序列化到结构体
func Bind(json string, v interface{}) error {
	var um fun.Unmarshall
	return um.UnmarshallBinding(json, v)
}

//序列化
func Marshall(v interface{}) (string, error) {
	var m fun.Marshall
	return m.MarshallJSON(v)
}
