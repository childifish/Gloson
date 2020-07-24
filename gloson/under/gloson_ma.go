package under

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type GlosonMa struct {
	Object interface{}
	ItemNum int
	TagMap map[int]Item
	Nesting bool
	Json string
}

type Item struct {
	Name string
	Value interface{}
}

func (glom *GlosonMa)StartMarshall(v interface{})error  {
	glom.Object = v

	err :=glom.ViewItem()
	if err != nil{
		return err
	}

	err = glom.Factory()
	if err != nil{
		return err
	}

	return nil
}

//找到key和Field总数
func (glom *GlosonMa)ViewItem()error  {
	typ := reflect.TypeOf(glom.Object).Elem()
	fNum := typ.NumField()
	glom.ItemNum = fNum
	fmt.Println("num",fNum)
	tagMap := make(map[int]Item)
	for i := 0; i < fNum; i++ {
		key := typ.Field(i).Tag.Get("json")
		if key == ""{
			return errors.New("void tag")
		}
		t :=reflect.ValueOf(glom.Object)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		value := t.Field(i).Interface()
			tagMap[i] =  Item{
			Name:  key,
			Value: value,
		}
	}
	glom.TagMap = tagMap
	fmt.Println("现在的map\n", glom.TagMap)
	return nil
}


func (glom *GlosonMa)Factory()error  {
	json := "{\n"
	//循环ItemNum int - 1次
	for i := 0; i< glom.ItemNum - 1;i++  {
		tag := glom.TagMap[i].Name
		json += "    \"" + tag + "\":"
		json += Value2String(glom.TagMap[i].Value) + ",\n"
	}
	tag := glom.TagMap[glom.ItemNum-1].Name
	json += "    \"" + tag + "\":"
	json += Value2String(glom.TagMap[glom.ItemNum-1].Value) + "\n}"
	glom.Json = json
	return nil
}

func Value2String(v interface{})string  {
	switch v.(type) {
	default:
		return ""
	case string:
		r  := "\"" + v.(string) + "\""
		return r
	case int:
		r := strconv.Itoa(v.(int))
		return r
	case bool:
		r := v.(bool)
		if r {
			return "true"
		}else {
			return "false"
		}
	case float64:
		b := v.(float64)
		r := strconv.FormatFloat(b, 'E', -1, 64)
		return r

	}
}