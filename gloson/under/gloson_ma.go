package under

import (
	"errors"
	"reflect"
	"strconv"
)

//序列化
type GlosonMa struct {
	Object  interface{}  //传入的结构体
	ItemNum int          //结构体field数量
	TagMap  map[int]Item //key--位置 value--信息
	Nesting []int        //位置
	Json    string       //完成后的json string
}

//key--位置 value--信息
type Item struct {
	Name  string
	Value interface{}
}

func (glom *GlosonMa) StartMarshall(v interface{}) error {
	glom.Object = v
	err := glom.ViewItem()
	if err != nil {
		return err
	}
	//中间有内嵌结构体
	if glom.Nesting != nil {

	}
	err = glom.Factory()
	if err != nil {
		return err
	}

	return nil
}

//找到key和Field总数
func (glom *GlosonMa) ViewItem() error {
	typ := reflect.TypeOf(glom.Object).Elem()
	fNum := typ.NumField()
	glom.ItemNum = fNum
	tagMap := make(map[int]Item)
	for i := 0; i < fNum; i++ {
		key := typ.Field(i).Tag.Get("json")
		if key == "" {
			return errors.New("void tag")
		}
		t := reflect.ValueOf(glom.Object)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Field(i).Kind() == reflect.Struct {
			glom.Nesting = append(glom.Nesting, i)
		}
		value := t.Field(i).Interface()
		tagMap[i] = Item{
			Name:  key,
			Value: value,
		}
	}
	glom.TagMap = tagMap
	return nil
}

func (glom *GlosonMa) Factory() error {
	json := "{\n"
	//循环ItemNum int - 1次
	for i := 0; i < glom.ItemNum-1; i++ {
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

//最基础的写入
func Value2String(v interface{}) string {
	switch v.(type) {
	default:
		return ""
	case string:
		r := "\"" + v.(string) + "\""
		return r
	case int:
		r := strconv.Itoa(v.(int))
		return r
	case bool:
		r := v.(bool)
		if r {
			return "true"
		} else {
			return "false"
		}
	case float64:
		b := v.(float64)
		r := strconv.FormatFloat(b, 'f', -1, 64)
		return r
	case []int:
		arr := v.([]int)
		str := "["
		l := len(arr)
		for i := 0; i < l-1; i++ {
			str += Value2String(arr[i]) + ", "
		}
		str += Value2String(arr[l-1]) + "]"
		return str
	case []float64:
		arr := v.([]float64)
		str := "["
		l := len(arr)
		for i := 0; i < l-1; i++ {
			str += Value2String(arr[i]) + ", "
		}
		str += Value2String(arr[l-1]) + "]"
		return str
	case []string:
		arr := v.([]string)
		str := "["
		l := len(arr)
		for i := 0; i < l-1; i++ {
			str += Value2String(arr[i]) + ", "
		}
		str += Value2String(arr[l-1]) + "]"
		return str

	}
}
