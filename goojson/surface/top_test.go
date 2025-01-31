package surface

import (
	"encoding/json"
	"fmt"
	"goojson/under"
	"reflect"
	"testing"
)

//简单json
func TestUnmarshall1(t *testing.T) {
	var un Unmarshall
	s := "{\n    \"data\": {\n        \"goods\": null\n    },\n    \"info\": \"success\",\n    \"status\": 200\n}"
	r, _ := un.Unmarshall().Map(s)
	for k, v := range r {
		fmt.Println(k, v)
	}
}

//较复杂json
func TestUnmarshall2(t *testing.T) {
	var un Unmarshall
	s := "{\n    \"info\": \"success\",\n    \"item\": {\n        \"1\": {\n            \"ID\": 1,\n            \"Name\": \"EHLVEEXN\",\n            \"Total\": 192,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"2\": {\n            \"ID\": 2,\n            \"Name\": \"CGSRGMVC\",\n            \"Total\": 253,\n            \"Left\": 147,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"3\": {\n            \"ID\": 3,\n            \"Name\": \"IXNBFIET\",\n            \"Total\": 275,\n            \"Left\": 145,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"4\": {\n            \"ID\": 4,\n            \"Name\": \"CXSWQGAL\",\n            \"Total\": 346,\n            \"Left\": 226,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"5\": {\n            \"ID\": 5,\n            \"Name\": \"CWHEGWDZ\",\n            \"Total\": 50,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"6\": {\n            \"ID\": 6,\n            \"Name\": \"SDFUNJFG\",\n            \"Total\": 168,\n            \"Left\": 52,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"7\": {\n            \"ID\": 7,\n            \"Name\": \"CNKLUXFR\",\n            \"Total\": 461,\n            \"Left\": 339,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"8\": {\n            \"ID\": 8,\n            \"Name\": \"FZQDICTN\",\n            \"Total\": 390,\n            \"Left\": 268,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"9\": {\n            \"ID\": 9,\n            \"Name\": \"WZWROCQQ\",\n            \"Total\": 298,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"10\": {\n            \"ID\": 10,\n            \"Name\": \"OJKEAGFX\",\n            \"Total\": 138,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"11\": {\n            \"ID\": 11,\n            \"Name\": \"UPEVWGDN\",\n            \"Total\": 307,\n            \"Left\": 194,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"12\": {\n            \"ID\": 12,\n            \"Name\": \"XQIKIVDB\",\n            \"Total\": 65,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"13\": {\n            \"ID\": 13,\n            \"Name\": \"DKNBZTBL\",\n            \"Total\": 148,\n            \"Left\": 36,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"14\": {\n            \"ID\": 14,\n            \"Name\": \"VUCNVUIO\",\n            \"Total\": 156,\n            \"Left\": 84,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"15\": {\n            \"ID\": 15,\n            \"Name\": \"ONNHPSLS\",\n            \"Total\": 69,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"16\": {\n            \"ID\": 16,\n            \"Name\": \"QBQKAUGX\",\n            \"Total\": 161,\n            \"Left\": 69,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"17\": {\n            \"ID\": 17,\n            \"Name\": \"EZKMKDIS\",\n            \"Total\": 297,\n            \"Left\": 175,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"18\": {\n            \"ID\": 18,\n            \"Name\": \"KSOKSBZC\",\n            \"Total\": 57,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"19\": {\n            \"ID\": 19,\n            \"Name\": \"YBUOXUPN\",\n            \"Total\": 200,\n            \"Left\": 62,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"20\": {\n            \"ID\": 20,\n            \"Name\": \"DMMPAORN\",\n            \"Total\": 288,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"21\": {\n            \"ID\": 21,\n            \"Name\": \"SZBYCNYW\",\n            \"Total\": 124,\n            \"Left\": 124,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"22\": {\n            \"ID\": 22,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"23\": {\n            \"ID\": 23,\n            \"Name\": \"tttttedt22\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"24\": {\n            \"ID\": 24,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"25\": {\n            \"ID\": 25,\n            \"Name\": \"tedt1\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"26\": {\n            \"ID\": 26,\n            \"Name\": \"\",\n            \"Total\": 0,\n            \"Left\": 0,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        }\n    },\n    \"status\": 200\n}"
	r, _ := un.Unmarshall().Map(s)
	for k, v := range r {
		va, ok := v.(map[string]interface{})
		if ok {
			for k, v := range va {
				fmt.Println(k, v)
			}
		}
		fmt.Println(k, v)
	}
}

//格式有问题的json
func TestUnmarshall3(t *testing.T) {
	var un Unmarshall
	s := "\n    \"info\": \"success\",\n    \"item\": {\n        \"1\": {\n            \"ID\": 1,\n            \"Name\": \"EHLVEEXN\",\n            \"Total\": 192,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"2\": {\n            \"ID\": 2,\n            \"Name\": \"CGSRGMVC\",\n            \"Total\": 253,\n            \"Left\": 147,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"3\": {\n            \"ID\": 3,\n            \"Name\": \"IXNBFIET\",\n            \"Total\": 275,\n            \"Left\": 145,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"4\": {\n            \"ID\": 4,\n            \"Name\": \"CXSWQGAL\",\n            \"Total\": 346,\n            \"Left\": 226,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"5\": {\n            \"ID\": 5,\n            \"Name\": \"CWHEGWDZ\",\n            \"Total\": 50,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"6\": {\n            \"ID\": 6,\n            \"Name\": \"SDFUNJFG\",\n            \"Total\": 168,\n            \"Left\": 52,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"7\": {\n            \"ID\": 7,\n            \"Name\": \"CNKLUXFR\",\n            \"Total\": 461,\n            \"Left\": 339,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"8\": {\n            \"ID\": 8,\n            \"Name\": \"FZQDICTN\",\n            \"Total\": 390,\n            \"Left\": 268,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"9\": {\n            \"ID\": 9,\n            \"Name\": \"WZWROCQQ\",\n            \"Total\": 298,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"10\": {\n            \"ID\": 10,\n            \"Name\": \"OJKEAGFX\",\n            \"Total\": 138,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"11\": {\n            \"ID\": 11,\n            \"Name\": \"UPEVWGDN\",\n            \"Total\": 307,\n            \"Left\": 194,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"12\": {\n            \"ID\": 12,\n            \"Name\": \"XQIKIVDB\",\n            \"Total\": 65,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"13\": {\n            \"ID\": 13,\n            \"Name\": \"DKNBZTBL\",\n            \"Total\": 148,\n            \"Left\": 36,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"14\": {\n            \"ID\": 14,\n            \"Name\": \"VUCNVUIO\",\n            \"Total\": 156,\n            \"Left\": 84,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"15\": {\n            \"ID\": 15,\n            \"Name\": \"ONNHPSLS\",\n            \"Total\": 69,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"16\": {\n            \"ID\": 16,\n            \"Name\": \"QBQKAUGX\",\n            \"Total\": 161,\n            \"Left\": 69,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"17\": {\n            \"ID\": 17,\n            \"Name\": \"EZKMKDIS\",\n            \"Total\": 297,\n            \"Left\": 175,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"18\": {\n            \"ID\": 18,\n            \"Name\": \"KSOKSBZC\",\n            \"Total\": 57,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"19\": {\n            \"ID\": 19,\n            \"Name\": \"YBUOXUPN\",\n            \"Total\": 200,\n            \"Left\": 62,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"20\": {\n            \"ID\": 20,\n            \"Name\": \"DMMPAORN\",\n            \"Total\": 288,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"21\": {\n            \"ID\": 21,\n            \"Name\": \"SZBYCNYW\",\n            \"Total\": 124,\n            \"Left\": 124,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"22\": {\n            \"ID\": 22,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"23\": {\n            \"ID\": 23,\n            \"Name\": \"tttttedt22\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"24\": {\n            \"ID\": 24,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"25\": {\n            \"ID\": 25,\n            \"Name\": \"tedt1\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"26\": {\n            \"ID\": 26,\n            \"Name\": \"\",\n            \"Total\": 0,\n            \"Left\": 0,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        }\n    },\n    \"status\": 200\n}"
	r, err := un.Unmarshall().Map(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range r {
		fmt.Println(k, v)
	}
}

// todo:带数组的json--有问题1
func TestUnmarshall4(t *testing.T) {
	var un Unmarshall
	s := "{ \"people\": [\n  { \"firstName\": \"Brett\", \"lastName\":\"McLaughlin\", \"email\": \"brett@newInstance.com\" },\n  { \"firstName\": \"Jason\", \"lastName\":\"Hunter\", \"email\": \"jason@servlets.com\" },\n  { \"firstName\": \"Elliotte\", \"lastName\":\"Harold\", \"email\": \"elharo@macfaq.com\" }\n]}"
	r, _ := un.Unmarshall().Map(s)
	Recursion(r)
}

// todo:带数组的json--有问题2
func TestUnmarshall5(t *testing.T) {
	var un Unmarshall
	s := "{\n\"name\":\"网站\",\n\"num\":3,\n\"sites\":[1,2,3]\n}"
	r, _ := un.Unmarshall().Map(s)
	Recursion(r)
}

//较复杂json--map嵌套
func TestUnmarshall6(t *testing.T) {
	var un Unmarshall
	s := "{\n    \"info\": \"success\",\n    \"item\": {\n        \"1\": {\n            \"ID\": 1,\n            \"Name\": \"EHLVEEXN\",\n            \"Total\": 192,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"2\": {\n            \"ID\": 2,\n            \"Name\": \"CGSRGMVC\",\n            \"Total\": 253,\n            \"Left\": 147,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"3\": {\n            \"ID\": 3,\n            \"Name\": \"IXNBFIET\",\n            \"Total\": 275,\n            \"Left\": 145,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"4\": {\n            \"ID\": 4,\n            \"Name\": \"CXSWQGAL\",\n            \"Total\": 346,\n            \"Left\": 226,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"5\": {\n            \"ID\": 5,\n            \"Name\": \"CWHEGWDZ\",\n            \"Total\": 50,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"6\": {\n            \"ID\": 6,\n            \"Name\": \"SDFUNJFG\",\n            \"Total\": 168,\n            \"Left\": 52,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"7\": {\n            \"ID\": 7,\n            \"Name\": \"CNKLUXFR\",\n            \"Total\": 461,\n            \"Left\": 339,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"8\": {\n            \"ID\": 8,\n            \"Name\": \"FZQDICTN\",\n            \"Total\": 390,\n            \"Left\": 268,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"9\": {\n            \"ID\": 9,\n            \"Name\": \"WZWROCQQ\",\n            \"Total\": 298,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"10\": {\n            \"ID\": 10,\n            \"Name\": \"OJKEAGFX\",\n            \"Total\": 138,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"11\": {\n            \"ID\": 11,\n            \"Name\": \"UPEVWGDN\",\n            \"Total\": 307,\n            \"Left\": 194,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"12\": {\n            \"ID\": 12,\n            \"Name\": \"XQIKIVDB\",\n            \"Total\": 65,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"13\": {\n            \"ID\": 13,\n            \"Name\": \"DKNBZTBL\",\n            \"Total\": 148,\n            \"Left\": 36,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"14\": {\n            \"ID\": 14,\n            \"Name\": \"VUCNVUIO\",\n            \"Total\": 156,\n            \"Left\": 84,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"15\": {\n            \"ID\": 15,\n            \"Name\": \"ONNHPSLS\",\n            \"Total\": 69,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"16\": {\n            \"ID\": 16,\n            \"Name\": \"QBQKAUGX\",\n            \"Total\": 161,\n            \"Left\": 69,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"17\": {\n            \"ID\": 17,\n            \"Name\": \"EZKMKDIS\",\n            \"Total\": 297,\n            \"Left\": 175,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"18\": {\n            \"ID\": 18,\n            \"Name\": \"KSOKSBZC\",\n            \"Total\": 57,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"19\": {\n            \"ID\": 19,\n            \"Name\": \"YBUOXUPN\",\n            \"Total\": 200,\n            \"Left\": 62,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"20\": {\n            \"ID\": 20,\n            \"Name\": \"DMMPAORN\",\n            \"Total\": 288,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"21\": {\n            \"ID\": 21,\n            \"Name\": \"SZBYCNYW\",\n            \"Total\": 124,\n            \"Left\": 124,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"22\": {\n            \"ID\": 22,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"23\": {\n            \"ID\": 23,\n            \"Name\": \"tttttedt22\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"24\": {\n            \"ID\": 24,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"25\": {\n            \"ID\": 25,\n            \"Name\": \"tedt1\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"26\": {\n            \"ID\": 26,\n            \"Name\": \"\",\n            \"Total\": 0,\n            \"Left\": 0,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        }\n    },\n    \"status\": 200\n}"
	r, _ := un.Unmarshall().Map(s)
	Recursion(r)
}

//特复杂的json
func TestUnmarshall7(t *testing.T) {
	var un Unmarshall
	s := "{\n\"glossary\": {\n\"title\": \"example glossary\",\n\"GlossDiv\": {\n\"title\": \"S\",\n\"GlossList\": {\n\"GlossEntry\": {\n\"ID\": \"SGML\",\n\"SortAs\": \"SGML\",\n\"GlossTerm\": \"Standard Generalized Markup Language\",\n\"Acronym\": \"SGML\",\n\"Abbrev\": \"ISO 8879:1986\",\n\"GlossDef\": {\n\"para\": \"A meta-markup language, used to create markup languages such as DocBook.\",\n\"GlossSeeAlso\": [\n\"GML\",\n\"XML\"\n]\n},\n\"GlossSee\": \"markup\"\n}\n}\n}\n}\n}"
	r, _ := un.Unmarshall().Map(s)
	Recursion(r)
}

//递归打印--测试用
func Recursion(input map[string]interface{}) {
	for k, v := range input {
		fmt.Println("key", k, "value", v, "type", reflect.TypeOf(v))
		value := v.(under.JsonKV)
		if value.WithinType == "option" {
			mapIn, ok := value.Within.(map[string]interface{})
			if ok {
				Recursion(mapIn)
			}
		}
		if value.WithinType == "array" {
			fmt.Println("array的遍历是坏的")
			mapIn2, ok := value.Within.([]interface{})
			fmt.Println("key", k, "value", v, "type", reflect.TypeOf(value.Within))
			if ok {
				for k, v := range mapIn2 {
					fmt.Println("keyin", k, "valuein", v, "typein", reflect.TypeOf(value.Within))
				}
			}
		}
	}
	return
}

//string&int 测试
func TestUnmarshall_Bind(t *testing.T) {
	var un Unmarshall
	s := "{\n    \"data\": \"goods\",\n    \"info\": \"success\",\n    \"status\": 200\n}"
	type TES struct {
		Info   string `json:"info"`
		Data   string `json:"data"`
		Status int    `json:"status"`
	}
	a := &TES{}
	un.Bind(s, a)
	fmt.Println(a)
}

//基础测试--原生
func TestUnmarshall_Bind2(t *testing.T) {
	//var un Unmarshall
	var tc under.TypeChanger
	s := "{\n    \"data\": \"goods\",\n    \"info\": \"success\",\n    \"status\": 200\n}"
	type TES struct {
		Info   string `json:"info"`
		Data   string `json:"data"`
		Status int    `json:"status"`
	}
	a := &TES{}
	json.Unmarshal(tc.Str2bytes(s), a)
	fmt.Println(a)
}

//嵌套结构体--不行
func TestUnmarshall_Bind3(t *testing.T) {
	var un Unmarshall
	s := "{\n    \"data\": {\n        \"goods\": 123145.00123\n    },\n    \"info\": \"success\",\n    \"status\": 200\n}"
	type Dat struct {
		Good int `json:"good"`
	}
	type TES struct {
		Info   string `json:"info"`
		Data   Dat    `json:"data"`
		Status int    `json:"status"`
	}
	a := &TES{}
	un.Bind(s, a)
	fmt.Println(a)
}

//好像官方的也不行
func TestUnmarshall_Bind4(t *testing.T) {
	var tc under.TypeChanger
	s := "{\n    \"data\": {\n        \"goods\": 1\n    },\n    \"info\": \"success\",\n    \"status\": 200\n}"
	type Data struct {
		Good int `json:"good"`
	}
	type TES struct {
		Info   string `json:"info"`
		Data   Data   `json:"data"`
		Status int    `json:"status"`
	}
	a := &TES{}
	err := json.Unmarshal(tc.Str2bytes(s), a)
	fmt.Println(err)
	fmt.Println(a)
}

//浮点数
func TestUnmarshall_Bind5(t *testing.T) {
	//var un Unmarshall
	var tc under.TypeChanger
	s := "{\n    \"data\": 110.2137162,\n    \"info\": \"success\",\n    \"status\": 200\n}"
	type TES struct {
		Info   string  `json:"info"`
		Data   float64 `json:"data"`
		Status int     `json:"status"`
	}
	a := &TES{}
	json.Unmarshal(tc.Str2bytes(s), a)
	fmt.Println(a)
}

//float64&int
func TestUnmarshall_Bind6(t *testing.T) {
	var un Unmarshall
	s := "{\n    \"data\": 113.12312445,\n    \"info\": \"success\",\n    \"status\": 200\n}"
	type TES struct {
		Info   string  `json:"info"`
		Data   float64 `json:"data"`
		Status int     `json:"status"`
	}
	a := &TES{}
	un.Bind(s, a)
	fmt.Println(a)
}

//longInt&bool
func TestUnmarshall_Bind7(t *testing.T) {
	var un Unmarshall
	s := "{\n    \"data\": 1131231244521342,\n    \"info\": false,\n    \"status\": \"null\"\n}"
	type TES struct {
		Info   bool   `json:"info"`
		Data   int64  `json:"data"`
		Status string `json:"status"`
	}
	a := &TES{}
	un.Bind(s, a)
	fmt.Println(a)
}

//int数组
func TestUnmarshall_Bind8(t *testing.T) {
	var un Unmarshall
	s := "{\n\"name\":\"网站\",\n\"num\":3,\n\"sites\":[1,2,3]\n}"
	type AutoGenerated struct {
		Name  string  `json:"name"`
		Num   int     `json:"num"`
		Sites []int64 `json:"sites"`
	}
	a := &AutoGenerated{}
	un.Bind(s, a)
	fmt.Println(a)
}

//string数组
func TestUnmarshall_Bind9(t *testing.T) {
	var un Unmarshall
	s := "{\n\"name\":\"网站\",\n\"num\":3,\n\"sites\":[\"1\",\"2\",\"3\"]\n}"
	type AutoGenerated struct {
		Name  string   `json:"name"`
		Num   int      `json:"num"`
		Sites []string `json:"sites"`
	}
	a := &AutoGenerated{}
	un.Bind(s, a)
	fmt.Println(a)
}

//float数组  完美
func TestUnmarshall_Bind10(t *testing.T) {
	var un Unmarshall
	s := "{\n\"name\":\"网站\",\n\"num\":3,\n\"sites\":[1.982,2.127361,3.123632]\n}"
	type AutoGenerated struct {
		Name  string    `json:"name"`
		Num   int       `json:"num"`
		Sites []float64 `json:"sites"`
	}
	a := &AutoGenerated{}
	un.Bind(s, a)
	fmt.Println(a)
}

//特复杂的究极套娃--大失败
func TestUnmarshall_Bind11(t *testing.T) {
	var un Unmarshall
	s := "{\n\"glossary\": {\n\"title\": \"example glossary\",\n\"GlossDiv\": {\n\"title\": \"S\",\n\"GlossList\": {\n\"GlossEntry\": {\n\"ID\": \"SGML\",\n\"SortAs\": \"SGML\",\n\"GlossTerm\": \"Standard Generalized Markup Language\",\n\"Acronym\": \"SGML\",\n\"Abbrev\": \"ISO 8879:1986\",\n\"GlossDef\": {\n\"para\": \"A meta-markup language, used to create markup languages such as DocBook.\",\n\"GlossSeeAlso\": [\n\"GML\",\n\"XML\"\n]\n},\n\"GlossSee\": \"markup\"\n}\n}\n}\n}\n}"
	type GlossDef struct {
		Para         string   `json:"para"`
		GlossSeeAlso []string `json:"GlossSeeAlso"`
	}
	type GlossEntry struct {
		ID        string   `json:"ID"`
		SortAs    string   `json:"SortAs"`
		GlossTerm string   `json:"GlossTerm"`
		Acronym   string   `json:"Acronym"`
		Abbrev    string   `json:"Abbrev"`
		GlossDef  GlossDef `json:"GlossDef"`
		GlossSee  string   `json:"GlossSee"`
	}
	type GlossList struct {
		GlossEntry GlossEntry `json:"GlossEntry"`
	}
	type GlossDiv struct {
		Title     string    `json:"title"`
		GlossList GlossList `json:"GlossList"`
	}
	type Glossary struct {
		Title    string   `json:"title"`
		GlossDiv GlossDiv `json:"GlossDiv"`
	}
	type AutoGenerated struct {
		Glossary Glossary `json:"glossary"`
	}
	a := &AutoGenerated{}
	un.Bind(s, a)
	fmt.Println(a)
}

//测试原生特复杂的究极套娃
func TestUnmarshall_Bind12(t *testing.T) {
	var tc under.TypeChanger
	s := "{\n\"glossary\": {\n\"title\": \"example glossary\",\n\"GlossDiv\": {\n\"title\": \"S\",\n\"GlossList\": {\n\"GlossEntry\": {\n\"ID\": \"SGML\",\n\"SortAs\": \"SGML\",\n\"GlossTerm\": \"Standard Generalized Markup Language\",\n\"Acronym\": \"SGML\",\n\"Abbrev\": \"ISO 8879:1986\",\n\"GlossDef\": {\n\"para\": \"A meta-markup language, used to create markup languages such as DocBook.\",\n\"GlossSeeAlso\": [\n\"GML\",\n\"XML\"\n]\n},\n\"GlossSee\": \"markup\"\n}\n}\n}\n}\n}"
	type GlossDef struct {
		Para         string   `json:"para"`
		GlossSeeAlso []string `json:"GlossSeeAlso"`
	}
	type GlossEntry struct {
		ID        string   `json:"ID"`
		SortAs    string   `json:"SortAs"`
		GlossTerm string   `json:"GlossTerm"`
		Acronym   string   `json:"Acronym"`
		Abbrev    string   `json:"Abbrev"`
		GlossDef  GlossDef `json:"GlossDef"`
		GlossSee  string   `json:"GlossSee"`
	}
	type GlossList struct {
		GlossEntry GlossEntry `json:"GlossEntry"`
	}
	type GlossDiv struct {
		Title     string    `json:"title"`
		GlossList GlossList `json:"GlossList"`
	}
	type Glossary struct {
		Title    string   `json:"title"`
		GlossDiv GlossDiv `json:"GlossDiv"`
	}
	type AutoGenerated struct {
		Glossary Glossary `json:"glossary"`
	}
	a := &AutoGenerated{}
	err := json.Unmarshal(tc.Str2bytes(s), a)
	fmt.Println(err)
	fmt.Println(a)
}

//解决套娃
func TestUnmarshall_Bind13(t *testing.T) {
	var un Unmarshall
	s := "{\n\"GlossDef\": {\n\"para\": \"A meta-markup language, used to create markup languages such as DocBook.\",\n\"GlossSeeAlso\": [\n\"GML\",\n\"XML\"\n]\n},\n\"GlossSee\": \"markup\"\n}"
	type GlossDef struct {
		Para         string   `json:"para"`
		GlossSeeAlso []string `json:"GlossSeeAlso"`
	}
	type AutoGenerated struct {
		GlossDef GlossDef `json:"GlossDef"`
		GlossSee string   `json:"GlossSee"`
	}
	a := &AutoGenerated{}
	un.Bind(s, a)
	fmt.Println(a)
}
