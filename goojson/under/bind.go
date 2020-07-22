package under

import (
	"fmt"
	"reflect"
)

//一开始是先确定结构体的tag，再从map里找对应的值
//可以先递归遍历map再find key
func RecursionBinding(input map[string]interface{}, v interface{}, rv reflect.Value) {
	typ := reflect.TypeOf(v).Elem()
	fNum := typ.NumField()
	var keySlice map[int]string
	keySlice = make(map[int]string)
	for i := 0; i < fNum; i++ {
		key := typ.Field(i).Tag.Get("json")
		keySlice[i] = key
	}
	RecursionMap(input, rv, keySlice)
}

func RecursionMap(input map[string]interface{}, rv reflect.Value, key map[int]string) {
	for k, v := range input {
		fmt.Println("key", k, "value", v, "type", reflect.TypeOf(v))
		value := v.(JsonKV)
		if value.WithinType == "option" {
			mapIn, ok := value.Within.(map[string]interface{})
			if ok {
				RecursionMap(mapIn, rv, key)
			}
		}
		if value.WithinType == "array" {
			fmt.Println("array的遍历是坏的")
			mapIn2, ok := value.Within.([]interface{})
			fmt.Println("key", k, "value", v, "type", reflect.TypeOf(value.Within))
			if ok {
				for k, v := range mapIn2 {
					fmt.Println("keyin", k, "valuein", v, "typein", reflect.TypeOf(value.Within))
				}
			}
		}
		for i, val := range key {
			k = ClearMark(k)
			fmt.Println(val)
			fmt.Println(k)
			if k == val {
				fmt.Println("111")
				WriteItem(value, i, rv)
			}
		}
	}
	return
}

func WriteItem(value JsonKV, i int, rv reflect.Value) {
	deep := value.Within
	v := rv.Elem().Field(i)
	ok := v.CanSet()
	if !ok {
		fmt.Println("不可以")
	}
	switch value.WithinType {
	case "string":
		new := deep.(string)
		v.SetString(new)
	case "int":
		new := deep.(int)
		r := int64(new)
		v.SetInt(r)
	}

}
