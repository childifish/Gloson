package function

import (
	"fmt"
	u "goojson/under"
)

var tc u.TypeChanger
var g u.GoojsonByte
var j u.JsonKV

type Factory struct {
}

func (f *Factory) Umarsh2Map(raw string) (map[string]interface{}, error) {
	var jsonSlice []u.JsonKV
	type fi map[string]interface{}
	finalMap := make(fi)
	//替换掉换行符与空格
	g = tc.Str2bytes(u.CleanSpace(raw))

	//判断是否是json
	err := u.CheckJson(g)

	if err != nil {
		return nil, err
	}
	//拆分
	jsonSlice = g.Go2KV()

	fmt.Println("结束了")

	finalMap, err = j.Writer(jsonSlice)

	return finalMap, err
}
