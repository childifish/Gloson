package under

import (
	"fmt"
	"reflect"
)

func RecursionBinding(input map[string]interface{}, v interface{}) {
	typ := reflect.TypeOf(v).Elem()
	fNum := typ.NumField()
	var keySlice map[int]string
	keySlice = make(map[int]string)
	for i := 0; i < fNum; i++ {
		key := typ.Field(i).Tag.Get("json")
		keySlice[i] = key
	}
	RecursionMap(input, v, keySlice)

}

func RecursionMap(input map[string]interface{}, inter interface{}, key map[int]string) {
	for k, v := range input {
		fmt.Println("key", k, "value", v, "type", reflect.TypeOf(v))
		value := v.(JsonKV)
		if value.WithinType == "option" {
			mapIn, ok := value.Within.(map[string]interface{})
			if ok {
				RecursionMap(mapIn, v, key)
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
				WriteItem(value, &inter, k, i)
			}
		}
	}
	return
}

func WriteItem(value JsonKV, inter interface{}, k string, i int) {
	deep := value.Within
	v := reflect.ValueOf(inter).Elem().FieldByName(k)
	ok := v.CanSet()
	if !ok {
		fmt.Println("不可以")
	}
	new := reflect.ValueOf(deep)
	fmt.Println("在写了", new)
	v.Set(new)
	fmt.Println("new v", v)
	fmt.Println("inter", inter)
}
