package under

import (
	"fmt"
	"reflect"
	"strconv"
)

//序列化
type GlosonMa struct {
	Object  interface{} //传入的结构体
	ItemNum int         //结构体field数量
	TagMap  []Item      //
	Json    string      //完成后的json string
}

//key--位置 value--信息
type Item struct {
	Tag      string
	Position []int
	Paste    []int
	Value    interface{}
	Type     reflect.Kind
}

func (glom *GlosonMa) StartMarshall(v interface{}) error {
	glom.Object = v
	glom.ViewItem([]int{})
	fmt.Println(glom)
	err := glom.Factory()
	if err != nil {
		return err
	}

	return nil
}

//找到key和Field总数
func (glom *GlosonMa) ViewItem(posBeyond []int) {
	var nowItem Item
	val := reflect.ValueOf(glom.Object)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := reflect.TypeOf(glom.Object)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	fNum := typ.NumField()
	for i := 0; i < fNum; i++ {
		nowPos := append(posBeyond, i)
		tag := typ.FieldByIndex(nowPos).Tag.Get("json")
		info := val.FieldByIndex(nowPos).Interface()
		kd := val.FieldByIndex(nowPos).Kind()
		nowItem = Item{
			Tag:      tag,
			Position: nowPos,
			Paste:    posBeyond,
			Value:    info,
			Type:     kd,
		}
		if kd == reflect.Struct {
			glom.TagMap = append(glom.TagMap, nowItem)
			glom.ViewItem(nowPos)
			continue
		}
		glom.TagMap = append(glom.TagMap, nowItem)
	}

}

func (glom *GlosonMa) Factory() error {
	json := "{\n"
	//循环ItemNum int - 1次
	for i := 0; i < len(glom.TagMap); i++ {

		if glom.TagMap[i].Type == reflect.Struct {
			//continue
		}
		json += glom.InWrite(i)
	}
	//json += glom.InWrite(len(glom.TagMap) - 1)
	glom.Json = json
	return nil
}

func (glom *GlosonMa) InWrite(i int) string {
	if i == len(glom.TagMap)-1 {
		tag := glom.TagMap[len(glom.TagMap)-1].Tag
		json := "    \"" + tag + "\":"
		json += Value2String(glom.TagMap[len(glom.TagMap)-1].Value) + "\n}"
		return json
	}
	tag := glom.TagMap[i].Tag
	json := "    \"" + tag + "\":"
	json += Value2String(glom.TagMap[i].Value) + ",\n"
	return json
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
