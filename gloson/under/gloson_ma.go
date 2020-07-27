package under

import (
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

func (glom *GlosonMa) StartMarshall(v interface{}) {
	glom.Object = v
	glom.ViewItem([]int{})
	glom.NewFactory()
}

//
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
	fNum := val.FieldByIndex(posBeyond).NumField()
	for i := 0; i < fNum; i++ {
		nowPos := append(posBeyond, i)
		//todo：这里后续可以优化
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
		} else {
			glom.TagMap = append(glom.TagMap, nowItem)
		}
	}

}

//没什么好说的
func (glom *GlosonMa) NewFactory() {
	json := "{\n"
	MAX := len(glom.TagMap)
	for i := 0; i < MAX; i++ {
		json += glom.InWrite(i)
	}
	for g != 0 {
		json += Space(g) + "},\n"
		g--
	}
	json += "}"
	glom.Json = json
}

//全局变量-标识已经写了的大括号
var g int

//
func (glom *GlosonMa) InWrite(i int) string {
	var json string
	if i >= 1 && i <= len(glom.TagMap)-1 {
		//"和上一个贴/依附的位置不相同", "而且上一个不是结构体"
		if (!reflect.DeepEqual(glom.TagMap[i].Paste, glom.TagMap[i-1].Paste)) && !(glom.TagMap[i-1].Type == reflect.Struct) {
			c := len(glom.TagMap[i-1].Position) - len(glom.TagMap[i].Position)
			for i2 := 0; i2 < c; i2++ {
				json += Space(g) + "},\n"
				g--
			}
		}
	}
	if glom.TagMap[i].Type == reflect.Struct {
		tag := glom.TagMap[i].Tag
		json += Space(len(glom.TagMap[i].Position)) + "\"" + tag + "\":{\n"
		g++
		return json
	}
	tag := glom.TagMap[i].Tag
	json += Space(len(glom.TagMap[i].Position)) + "\"" + tag + "\":"
	json += Value2String(glom.TagMap[i].Value) + ",\n"
	return json
}

//没用到
func (glom *GlosonMa) FinalWrite() string {
	tag := glom.TagMap[len(glom.TagMap)-1].Tag
	json := Space(len(glom.TagMap[len(glom.TagMap)-1].Position)) + "\"" + tag + "\":"
	json += Value2String(glom.TagMap[len(glom.TagMap)-1].Value) + "\n"
	json += "}"
	return json
}

//用来加空格
func Space(i int) string {
	var p string
	for j := 0; j < i; j++ {
		p += "    "
	}
	return p
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
