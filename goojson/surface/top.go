package surface

import (
	  "goojson/function"
)

var f function.Factory

//Unmarshall
type Unmarshall string

func (um *Unmarshall)Unmarshall()*Unmarshall  {
	return um
}

func (um *Unmarshall)Map(json string)(map[string]interface{},error)  {
	return f.Umarsh2Map(json)
}

