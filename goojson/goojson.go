package goojson

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

type GooJson struct {
	Json interface{}
	NotJson interface{}
	Err error
}

type JsonRaw struct {
	Key string
	Within interface{}
	WithinType string
}

type JsonSlice struct{
	GooSlice []JsonRaw
}

//json序列化
func (gj *GooJson)Marshall()  {

}

//高效str2bytes
func str2bytes(s string) []byte {
    x := (*[2]uintptr)(unsafe.Pointer(&s))
   h := [3]uintptr{x[0], x[1], x[1]}
    return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}

//json反序列化
//使用例：json,err := jsonRaw.Unmarshall(&struct)
func (gj *GooJson)Unmarshall(stc interface{})(GooJson,error)  {
	//断言
	rawBytes,err := WhetherString2Bytes(gj.Json)
	if err != nil{
		return GooJson{Err: err},err
	}

	//
	//RawBytes-->Chuck2JsonRaw-->chuck-->itemSlice-->JsonRawSlice
	rawSlice,err := RawBytes2JsonRaw(rawBytes)

	if err != nil{
		return GooJson{Err: err},err
	}

	//[]JsonRaw-->KeyMap(类型转换)
	KeyMap,err := JsonRaw2KeyMap(rawSlice)
	if err != nil{
		return GooJson{Err: err},err
	}

	fmt.Println(KeyMap)

	//KeyMap绑结构体
	stcAfter,err := BindStruct(KeyMap,stc)
	if err != nil{
		return GooJson{Err: err},err
	}

	//从
	return GooJson{NotJson: stcAfter},err
}

//string或者[]byte-->[]byte  空接口断言为
func WhetherString2Bytes(v interface{})([]byte,error)  {
	if b,ok:=v.([]byte);ok{
		s := bytes2str(b)
		return str2bytes(ClearSpace(s)),nil
	}else if s,ok2:=v.(string);ok2{
		return str2bytes(ClearSpace(s)),nil
	}
	return nil,errors.New("not string or bytes")
}

//将一个
func RawBytes2JsonRaw(rawBytes []byte)([]JsonRaw,error)  {
	if rawBytes[0]!=123||rawBytes[len(rawBytes)-1]!=125{
		return nil,errors.New("not correct json")
	}else {
		//有strings.Split()但还是想写个自己的
		//后来发现还真得自己写，不然不能处理结构体套娃
		items := RawByte2Slice(rawBytes)
		//fmt.Println(items)
		JsonRawSlice:= Item2JsonRaw(items)
		//fmt.Println(JsonRawSlice)
		return JsonRawSlice,nil
	}
}

func Item2JsonRaw(items []string)[]JsonRaw {
	return BreakItem(items)
}

func JsonRaw2KeyMap(jsrSlice []JsonRaw)(map[string]interface{},error)  {
	//fmt.Println("jsr",jsrSlice)
	var valueMap = make(map[string]interface{})
	for _,v := range jsrSlice{
		key := v.Key
		err := v.CheckType()
		if err!=nil{
			//fmt.Println(err)
			return nil,err
		}
		valueMap[key] = v
	}
	return valueMap,nil
}

func ClearSpace(s string)string  {
	re:=strings.ReplaceAll(s,"\n","")
	return strings.ReplaceAll(re," ","")
}

func BindStruct(keyMap map[string]interface{},stc interface{})(interface{},error)  {
	return stc,nil
}

