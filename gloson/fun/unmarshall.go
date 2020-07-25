package fun

import (
	"gloson/under"
)

type Unmarshall struct {
	Gloson under.GlosonUma
}

//反序列化--到map
func (um *Unmarshall) Unmarshall2Map(raw string) (map[string]interface{}, error) {
	err := um.Gloson.StartMap(raw)
	return um.Gloson.Map, err
}

//反序列化--到结构体
func (um *Unmarshall) UnmarshallBinding(raw string, v interface{}) error {
	err := um.Gloson.StartBinding(raw, v)
	return err
}
