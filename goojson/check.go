package goojson

import (
	"errors"
	"strconv"
)

func (jsr *JsonRaw)CheckType()error  {
	_,ok := jsr.Distributor()
	if !ok {
		return errors.New("TypeError")
	}
	return nil
}

//分配器：根据within类型的不同返回不同结果
func (jsr *JsonRaw)Distributor()(interface{},bool)  {
	//fmt.Println("distribute",jsr)
	switch jsr.WithinType {
	default:
		//fmt.Println("unknown",jsr)
		return jsr.Nil()
	case "option":
		return jsr.Option()
	case "number":
		return jsr.Num()
	case "bool":
		return  jsr.Bool()
	case "array":
		return jsr.Array()
	case "string":
		return jsr.String()
	case "null":
		return jsr.Nil()
	}

}


//通过空接口和断言实现类型转换

func (jsr *JsonRaw)Bool()(bool,bool)  {
	return true,true
}

func (jsr *JsonRaw)Int()(int,bool)  {
	i, err := strconv.Atoi(jsr.Within.(string))
	if err != nil {
		return 0,false
	}
	jsr.Within = i
	jsr.WithinType = "int"
	return i,true
}

func (jsr *JsonRaw)Int64()(int64,bool)  {
	i, err := strconv.ParseInt(jsr.Within.(string), 10, 64)
	if err != nil {
		return 0,false
	}
	jsr.Within = i
	jsr.WithinType = "int64"
	return i,true
}

func (jsr *JsonRaw)Float64()(float64,bool)  {
	i, err := strconv.ParseFloat(jsr.Within.(string), 64)
	if err != nil {
		return 0,false
	}
	jsr.Within = i
	jsr.WithinType = "float64"
	return i,true
}

func (jsr *JsonRaw)Num()(interface{},bool)  {
	var ok bool
	var returner interface{}
	numS,_ :=jsr.Within.(string)
	numBs := str2bytes(numS)
	length := len(numBs)
	if length > 11 {
		returner,ok = jsr.Int64()
	}else {
		returner,ok = jsr.Int()
	}
	for _,nums := range numBs{
		//带小数点
		if nums == 46 {
			if length > 8{
				returner,ok = jsr.Float64()
			}
		}
	}
	return returner,ok
}

func (jsr *JsonRaw)String()(string,bool)  {
	b,ok:=jsr.Within.(string)
	if ok{
		jsr.WithinType = "string"
		return b,ok
	}
	return "",ok
}

func (jsr *JsonRaw)Nil()(interface{},bool)  {
	jsr.WithinType = "nil"
	return nil,true
}

func (jsr *JsonRaw)Array()([]interface{},bool)  {
	return jsr.Within.([]interface{}),true
}

func (jsr *JsonRaw)Option()(map[string]interface{},bool)  {
	return jsr.Unmarshall()
}

func (jsr *JsonRaw)Unmarshall()(map[string]interface{},bool)  {
	//fmt.Println("jsr in umm",jsr.Within)
	rawBytes,err := WhetherString2Bytes(jsr.Within)
	if err != nil{
		return nil,false
	}
	//fmt.Println("rbbbbbbb",rawBytes)
	rawSlice,err := RawBytes2JsonRaw(rawBytes)
	if err != nil{
		return nil,false
	}
	//for ey,v := range rawSlice {
	//	fmt.Println("option",ey,v)
	//}
	kMap,err := JsonRaw2KeyMap(rawSlice)
	jsr.Within = kMap
	return kMap,true
}