package under

import (
	"fmt"
	"reflect"
)

//一开始是先确定结构体的tag，再从map里找对应的值
//可以先递归遍历map再find key
func RecursionBinding(input map[string]interface{}, v interface{}) {
	keyMap := FindKeyD(v)
	RecursionMap(input, v, keyMap)
}

func FindKeyD(v interface{}) map[int]string {
	typ := reflect.TypeOf(v).Elem()
	fNum := typ.NumField()
	var keyMap map[int]string
	keyMap = make(map[int]string)
	for i := 0; i < fNum; i++ {
		key := typ.Field(i).Tag.Get("json")
		keyMap[i] = key
	}
	fmt.Println("现在的KEY", keyMap)
	return keyMap
}

func RecursionMap(input map[string]interface{}, vin interface{}, key map[int]string) {
	rv := reflect.ValueOf(vin)
	for k, v := range input {
		fmt.Println("key", k, "value", v, "type", reflect.TypeOf(v))
		value := v.(JsonKV)
		if value.WithinType == "option" {
			typ := reflect.TypeOf(vin).Elem()
			fNum := typ.NumField()
			for i := 0; i < fNum; i++ {
				type1 := typ.Field(i).Type
				//if type1 == reflect.Struct {
				//
				//}
				fmt.Println("type1-----------", type1)
			}
		}
		fmt.Println("现在的KEY--else", key)
		for i, val := range key {
			k = cl.CleanMark(k)
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
		fmt.Println("不可以", v)
	}
	switch value.WithinType {
	case "string":
		new := deep.(string)
		v.SetString(new)
	case "int", "int64":
		new, ok := deep.(int)
		if ok {
			r := int64(new)
			v.SetInt(r)
		} else {
			new, _ := deep.(int64)
			v.SetInt(new)
		}
	case "float64":
		new := deep.(float64)
		v.SetFloat(new)
	case "bool":
		new := deep.(bool)
		v.SetBool(new)
	case "array":
		fmt.Println("我是数组", deep)
		in, tag := tc.Bytes2array(tc.Str2bytes(deep.(string)))
		switch tag {
		case "string":
			var vr reflect.Value
			e0 := make([]reflect.Value, 0)
			for _, value := range in {
				e0 = append(e0, reflect.ValueOf(value.(string)))
				vr = reflect.Append(v, e0...)
			}
			v.Set(vr)
		case "int64":
			var vr reflect.Value
			e0 := make([]reflect.Value, 0)
			for _, value := range in {
				e0 = append(e0, reflect.ValueOf(value.(int64)))
				vr = reflect.Append(v, e0...)
			}
			v.Set(vr)
		case "float64":
			var vr reflect.Value
			e0 := make([]reflect.Value, 0)
			for _, value := range in {
				fmt.Println("ref  ", reflect.TypeOf(value))
				e0 = append(e0, reflect.ValueOf(value.(float64)))

				vr = reflect.Append(v, e0...)

			}
			v.Set(vr)
		}
	case "option":
		fmt.Println("我是结构体", v)
		rv = reflect.Indirect(rv)
		for i := 0; i < rv.NumField(); i++ {
			fmt.Println("芜湖！！", rv)
			WriteOption(value.Within.(map[string]interface{}), i, rv.Field(i).Addr())
		}
	case "nil":
		fmt.Println("我是空的", v)
	}

}

func WriteOption(m map[string]interface{}, i int, rf reflect.Value) {
	fmt.Println("rf", &rf)
	fmt.Println(FindKeyUD(rf))
	fmt.Println("坠机", m, i, rf)
}

func FindKeyUD(v interface{}) map[int]string {
	typ := reflect.TypeOf(v)
	fNum := typ.NumField()
	var keyMap map[int]string
	keyMap = make(map[int]string)
	for i := 0; i < fNum; i++ {
		fmt.Println("i", typ.Field(i))
		key := typ.Field(i).Tag.Get("json")
		keyMap[i] = key
	}
	fmt.Println("现在的KEY", keyMap)
	return keyMap
}
