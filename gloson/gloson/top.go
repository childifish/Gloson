package gloson

import "gloson/fun"


func Map(json string)(map[string]interface{}, error)  {
	var um fun.Unmarshall
	return um.Unmarshall2Map(json)
}

func Bind(json string,v interface{}) error {
	var um fun.Unmarshall
	return um.UnmarshallBinding(json,v)
}

func Marshall(v interface{})(string,error)  {
	var m fun.Marshall
	return m.MarshallJSON(v)
}