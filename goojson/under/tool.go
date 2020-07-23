package under

import (
	"fmt"
	"strconv"
	"unsafe"
)

type TypeChanger struct {
}

func (tc *TypeChanger) Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func (tc *TypeChanger) Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func (tc *TypeChanger) Bytes2array(b []byte) ([]interface{}, string) {
	var tag string
	var returner []string
	var v []interface{}
	mark := 1
	flag := true
	tagflag := false
	for i, v := range b {
		fmt.Println(i, v)
		if v == 34 { //"
			flag = !flag
		}
		if v == 46 {
			if flag {
				if !tagflag {
					tag = "float64"
					fmt.Println("dffffffffff")
					tagflag = true
				}
			}
		}
		if v == 44 {
			if flag {
				//数字
				if ((b[i+1] >= 48) && (b[i+1] <= 57)) || (b[i+1] == 45) {
					if !tagflag {
						tag = "int64"
						fmt.Println("iiiiiiiii64")
					}
				}
				if b[i+1] == 34 {
					tag = "string"
					fmt.Println("sssssssss")
					tagflag = true
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
	fmt.Println("re", returner)
	return v, tag
}
