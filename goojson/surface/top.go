package surface

import (
	"goojson/function"
)

var f function.Factory

//Unmarshall
type Unmarshall struct {
}

func (um *Unmarshall) Unmarshall() *Unmarshall {
	return um
}

func (um *Unmarshall) Map(json string) (map[string]interface{}, error) {
	return f.Umarsh2Map(json)
}

func (um *Unmarshall) Bind(json string, v interface{}) error {
	return f.Umarsh2BindStruct(json, v)
}
