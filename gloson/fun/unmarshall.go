package fun

import (
	"gloson/under"
)

type Unmarshall struct {
	Gloson under.GlosonUma
}

func (um *Unmarshall)Unmarshall2Map(raw string) (map[string]interface{}, error) {
	err := um.Gloson.StartMap(raw)
	return um.Gloson.Map,err
}

func (um *Unmarshall)UnmarshallBinding(raw string,v interface{})error{
	err := um.Gloson.StartBinding(raw,v)
	return err
}
