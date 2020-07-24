package under

import (
	"errors"
	"fmt"
	"strconv"
)

type JsonKV struct {
	Key        string
	Within     interface{}
	WithinType string
}

func (jkv *JsonKV) WriteInMap(jsrSlice []JsonKV) (map[string]interface{}, error) {
	//fmt.Println("writer", jsrSlice)
	var valueMap = make(map[string]interface{})
	for _, v := range jsrSlice {
		key := v.Key
		err := v.CheckType()
		if err != nil {
			return nil, err
		}
		valueMap[key] = v
	}
	return valueMap, nil
}

func (jkv *JsonKV) CheckType() error {
	_, ok := jkv.Distributor()
	if !ok {
		return errors.New("TypeError")
	}
	return nil
}

func (jkv *JsonKV) Distributor() (interface{}, bool) {
	fmt.Println("distribute", jkv)
	switch jkv.WithinType {
	default:
		return jkv.Nil()
	case "option":
		return jkv.Option()
	case "number":
		return jkv.Num()
	case "true":
		return jkv.True()
	case "false":
		return jkv.False()
	case "array":
		arr, ok := jkv.Array()
		if ok {
			return arr, ok
		} else {
			//是简单array
			s := jkv.Within.(string)
			ar, _ := tc.Bytes2array(tc.Str2bytes(s))
			return ar, !ok
		}
	case "string":
		return jkv.String()
	case "null":
		return jkv.Nil()
	}

}

func (jkv *JsonKV) True() (bool, bool) {
	jkv.Within = true
	jkv.WithinType = "bool"
	return true, true
}

func (jkv *JsonKV) False() (bool, bool) {
	jkv.Within = false
	jkv.WithinType = "bool"
	return false, true
}

func (jkv *JsonKV) Int() (int, bool) {
	i, err := strconv.Atoi(jkv.Within.(string))
	if err != nil {
		return 0, false
	}
	jkv.Within = i
	jkv.WithinType = "int"
	return i, true
}

func (jkv *JsonKV) Int64() (int64, bool) {
	i, err := strconv.ParseInt(jkv.Within.(string), 10, 64)
	if err != nil {
		return 0, false
	}
	jkv.Within = i
	jkv.WithinType = "int64"
	return i, true
}

func (jkv *JsonKV) Float64() (float64, bool) {
	i, err := strconv.ParseFloat(jkv.Within.(string), 64)
	if err != nil {
		return 0, false
	}
	jkv.Within = i
	jkv.WithinType = "float64"
	return i, true
}

func (jkv *JsonKV) Num() (interface{}, bool) {
	var ok bool
	var returner interface{}
	numS, _ := jkv.Within.(string)
	numBs := tc.Str2bytes(numS)
	length := len(numBs)
	if length > 11 {
		returner, ok = jkv.Int64()
	} else {
		returner, ok = jkv.Int()
	}
	for _, nums := range numBs {
		//带小数点
		if nums == 46 {
			returner, ok = jkv.Float64()
		}
	}
	return returner, ok
}

func (jkv *JsonKV) String() (string, bool) {
	b, ok := jkv.Within.(string)
	if ok {
		return b, ok
	}
	return "", ok
}

func (jkv *JsonKV) Nil() (interface{}, bool) {
	jkv.WithinType = "nil"
	return nil, true
}

func (jkv *JsonKV) Array() ([]interface{}, bool) {
	fmt.Println("触发Array")
	fmt.Println("array", jkv.Within)
	var g GlosonByte
	var returner []interface{}
	g = tc.Str2bytes(cl.CleanSpace(jkv.Within.(string)))
	fmt.Println("我是g", tc.Bytes2str(g))
	arr := g.Go2Array()
	if arr == nil {
		fmt.Println("已经是最佳", g)
		return nil, false
	}
	for _, v := range arr {
		returner = append(returner, v)
	}
	jkv.Within = returner
	return returner, true
}

func (jkv *JsonKV) Option() (map[string]interface{}, bool) {
	return jkv.Unmarshall()
}

func (jkv *JsonKV) Unmarshall() (map[string]interface{}, bool) {
	var g GlosonByte
	g = tc.Str2bytes(cl.CleanSpace(jkv.Within.(string)))
	fmt.Println(tc.Bytes2str(g))
	if len(g) == 2 {
		jkv.WithinType = "nil"
		return nil, true
	}
	rawSlice := g.Go2KV()
	kMap, err := jkv.WriteInMap(rawSlice)
	if err != nil {
		return nil, false
	}
	jkv.Within = kMap
	for k, v := range kMap {
		fmt.Println("im kmap", k, v)
	}

	return kMap, true
}