package fun

import "gloson/under"

type Marshall struct {
	Gloson under.GlosonMa
}

//序列化
func (m *Marshall) MarshallJSON(v interface{}) string {
	m.Gloson.StartMarshall(v)
	return m.Gloson.Json
}
