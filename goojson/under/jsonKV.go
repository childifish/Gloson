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

//Writer将[]JsonKV写入Map，同时细分value的类型
//CheckType之后类型是确定的
func (jsr *JsonKV) Writer(jsrSlice []JsonKV) (map[string]interface{}, error) {
	fmt.Println("writer", jsrSlice)
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

//
func (jsr *JsonKV) CheckType() error {
	fmt.Println("触发Check")
	_, ok := jsr.Distributor()
	if !ok {
		return errors.New("TypeError")
	}
	return nil
}

//分配器：根据within类型的不同返回不同结果
func (jsr *JsonKV) Distributor() (interface{}, bool) {
	fmt.Println("distribute", jsr)
	switch jsr.WithinType {
	default:
		//fmt.Println("unknown",jsr)
		return jsr.Nil()
	case "option":
		return jsr.Option()
	case "number":
		return jsr.Num()
	case "true":
		return jsr.True()
	case "false":
		return jsr.False()
	case "array":
		fmt.Println("wozai")
		arr, ok := jsr.Array()
		if ok {
			return arr, ok
		} else {
			fmt.Println("1111")
			s := jsr.Within.(string)
			fmt.Println(s)
			ar, _ := tc.Bytes2array(tc.Str2bytes(s))
			fmt.Println(ar)
			return ar, !ok
		}
	case "string":
		return jsr.String()
	case "null":
		return jsr.Nil()
	}

}

//通过空接口和断言实现类型转换

func (jsr *JsonKV) True() (bool, bool) {
	jsr.Within = true
	jsr.WithinType = "bool"
	return true, true
}

func (jsr *JsonKV) False() (bool, bool) {
	jsr.Within = false
	jsr.WithinType = "bool"
	return false, true
}

func (jsr *JsonKV) Int() (int, bool) {
	i, err := strconv.Atoi(jsr.Within.(string))
	if err != nil {
		return 0, false
	}
	jsr.Within = i
	jsr.WithinType = "int"
	return i, true
}

func (jsr *JsonKV) Int64() (int64, bool) {
	i, err := strconv.ParseInt(jsr.Within.(string), 10, 64)
	if err != nil {
		return 0, false
	}
	jsr.Within = i
	jsr.WithinType = "int64"
	return i, true
}

func (jsr *JsonKV) Float64() (float64, bool) {
	i, err := strconv.ParseFloat(jsr.Within.(string), 64)
	if err != nil {
		return 0, false
	}
	jsr.Within = i
	jsr.WithinType = "float64"
	return i, true
}

func (jsr *JsonKV) Num() (interface{}, bool) {
	var ok bool
	var returner interface{}
	numS, _ := jsr.Within.(string)
	numBs := tc.Str2bytes(numS)
	length := len(numBs)
	if length > 11 {
		returner, ok = jsr.Int64()
	} else {
		returner, ok = jsr.Int()
	}
	for _, nums := range numBs {
		//带小数点
		if nums == 46 {
			returner, ok = jsr.Float64()
		}
	}
	return returner, ok
}

func (jsr *JsonKV) String() (string, bool) {
	b, ok := jsr.Within.(string)
	if ok {
		return b, ok
	}
	return "", ok
}

func (jsr *JsonKV) Nil() (interface{}, bool) {
	jsr.WithinType = "nil"
	return nil, true
}

func (jsr *JsonKV) Array() ([]interface{}, bool) {
	fmt.Println("触发Array")
	fmt.Println("array", jsr.Within)
	var g GoojsonByte
	var returner []interface{}
	g = tc.Str2bytes(CleanSpace(jsr.Within.(string)))
	fmt.Println("我是g", tc.Bytes2str(g))
	arr := g.Go2Array()
	if arr == nil {
		fmt.Println("已经是最佳", g)
		return nil, false
	}
	for _, v := range arr {
		returner = append(returner, v)
	}
	jsr.Within = returner
	return returner, true
}

func (jsr *JsonKV) Option() (map[string]interface{}, bool) {
	return jsr.Unmarshall()
}

func (jsr *JsonKV) Unmarshall() (map[string]interface{}, bool) {
	var g GoojsonByte
	g = tc.Str2bytes(CleanSpace(jsr.Within.(string)))
	fmt.Println(tc.Bytes2str(g))
	if len(g) == 2 {
		jsr.WithinType = "nil"
		return nil, true
	}
	rawSlice := g.Go2KV()
	kMap, err := jsr.Writer(rawSlice)
	if err != nil {
		return nil, false
	}
	jsr.Within = kMap
	for k, v := range kMap {
		fmt.Println("im kmap", k, v)
	}

	return kMap, true
}
