package under

import (
	"errors"
	//"fmt"
	"strconv"
	"strings"
	"unsafe"
)

type TypeChanger struct {
}
type Cleaner struct {
}

var tc TypeChanger
var cl Cleaner

//去除传入字符串里的空格和换行符
func (cl *Cleaner) CleanSpace(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), " ", "")
}

//检测是否是正确的Json
func (cl *Cleaner) CheckJson(rawBytes []byte) error {
	if len(rawBytes) != 0 {
		if rawBytes[0] != 123 || rawBytes[len(rawBytes)-1] != 125 {
			return errors.New("not correct json")
		}
		return nil
	} else {
		return errors.New("nil json")
	}
}

//去除Key里的引号
func (cl *Cleaner) CleanMark(s string) string {
	b := tc.Str2bytes(s)
	if b[0] == 34 {
		b = b[1 : len(b)-1]
		return tc.Bytes2str(b)
	}
	return tc.Bytes2str(b)
}

//高效字符串转Bytes
func (tc *TypeChanger) Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//高效Bytes数组转字符串
func (tc *TypeChanger) Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

//字符串数组转简单数组（没有嵌套结构）和数组储存类型（以一个string Tag表示）
func (tc *TypeChanger) Bytes2array(b []byte) ([]interface{}, string) {
	var tag string
	var returner []string
	var v []interface{}
	mark := 1
	flag := true
	tagFlag := false
	for i, v := range b {
		if v == 34 { //"
			flag = !flag
		}
		if v == 46 {
			if flag {
				if !tagFlag {
					tag = "float64"
					tagFlag = true
				}
			}
		}
		if v == 44 {
			if flag {
				//数字
				if ((b[i+1] >= 48) && (b[i+1] <= 57)) || (b[i+1] == 45) {
					if !tagFlag {
						tag = "int64"
					}
				}
				if b[i+1] == 34 {
					tag = "string"
					tagFlag = true
				}
				returner = append(returner, tc.Bytes2str(b[mark:i]))
				mark = i + 1
			}
		}
	}
	returner = append(returner, tc.Bytes2str(b[mark:len(b)-1]))
	switch tag {
	case "string":
		for _, value := range returner {
			v = append(v, value)
		}
	case "int64":
		for _, value := range returner {
			vii, _ := strconv.ParseInt(value, 10, 64)
			v = append(v, vii)
		}
	case "float64":
		for _, value := range returner {
			vii, _ := strconv.ParseFloat(value, 64)
			v = append(v, vii)
		}
	}
	return v, tag
}

//字符串，数字（以及他们的数组）转string
func (tc *TypeChanger) Value2String(v interface{}) string {
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
