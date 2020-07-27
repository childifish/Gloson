package under

import (
	"reflect"
)

//绑定 //NestAbility是里面是否有结构体
type BindingMsg struct {
	BindBefore  interface{}
	SurfaceKey  map[string]Field
	NestAbility bool
}

//类型和位置，其中位置是一个数组，从结构体反射获得
type Field struct {
	Position []int
	Type     reflect.Kind
}

//解析传入的结构体
func (bin *BindingMsg) InitBindingMsg(v interface{}) {
	bin.BindBefore = v
	bin.SurfaceKey = make(map[string]Field)
	bin.FindKeyD([]int{})

}

//解析结构里的全部tag
func (bin *BindingMsg) FindKeyD(posBeyond []int) {
	val := reflect.ValueOf(bin.BindBefore)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := reflect.TypeOf(bin.BindBefore)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	fNum := typ.NumField()
	var keyMap map[string]Field
	keyMap = make(map[string]Field)
	for i := 0; i < fNum; i++ {
		nowPos := append(posBeyond, i)
		key := typ.FieldByIndex(nowPos).Tag.Get("json")
		keyMap[key] = Field{
			Position: nowPos,
			Type:     val.FieldByIndex(nowPos).Kind(),
		}
	}
	for k, v := range keyMap {
		bin.SurfaceKey[k] = v
		if v.Type == reflect.Struct {
			var p []int
			for _, v := range v.Position {
				p = append(p, v)
			}
			bin.FindKeyD(p)
		}
	}
}

//找field
func Checkin(keymap map[string]Field, tag string) (Field, bool) {
	for k, _ := range keymap {
		if k == tag {
			return keymap[k], true
		}
	}
	return Field{}, false
}

//写入
func (bin *BindingMsg) WriteItem(value JsonKV, i []int) {
	deep := value.Within
	rv := reflect.ValueOf(bin.BindBefore)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	v := rv.FieldByIndex(i)
	switch value.WithinType {
	case "string":
		n := deep.(string)
		v.SetString(n)
	case "int", "int64":
		n, ok := deep.(int)
		if ok {
			r := int64(n)
			v.SetInt(r)
		} else {
			n2, _ := deep.(int64)
			v.SetInt(n2)
		}
	case "float64":
		n := deep.(float64)
		v.SetFloat(n)
	case "bool":
		n := deep.(bool)
		v.SetBool(n)
	case "array":
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
				e0 = append(e0, reflect.ValueOf(value.(float64)))
				vr = reflect.Append(v, e0...)
			}
			v.Set(vr)
		}
	case "option":
		var b BindingMsg
		b.InitBindingMsg(bin.BindBefore)
		mp := value.Within.(map[string]interface{})
		for _, j := range mp {
			kv := j.(JsonKV)
			tag := cl.CleanMark(kv.Key)
			field, _ := Checkin(b.SurfaceKey, tag)
			b.WriteItem(kv, field.Position)
		}
	case "nil":
	}
	delete(bin.SurfaceKey, cl.CleanMark(value.Key))
}
