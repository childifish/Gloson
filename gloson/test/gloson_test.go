package test

import (
	"encoding/json"
	"fmt"
	"gloson/gloson"
	"gloson/under"
	"reflect"
	"testing"
)

//递归打印map
func RecursionPrint(input map[string]interface{}) {
	for k, v := range input {
		fmt.Println("key", k, "value", v, "type", reflect.TypeOf(v))
		value := v.(under.JsonKV)
		if value.WithinType == "option" {
			mapIn, ok := value.Within.(map[string]interface{})
			if ok {
				RecursionPrint(mapIn)
			}
		}
	}
	return
}

//简单反序列化--map
func Test(t *testing.T) {
	s := "{\n    \"data\": {\n        \"goods\": null\n    },\n    \"info\": \"success\",\n    \"status\": 200\n}"
	map1, _ := gloson.Map(s)
	fmt.Println(map1)
}

//简单反序列化--struct
func Test2(t *testing.T) {
	s := "{\n\"name\":\"网站\",\n\"num\":3,\n\"sites\":[1,2,3]\n}"
	type AutoGenerated struct {
		Name  string  `json:"name"`
		Num   int     `json:"num"`
		Sites []int64 `json:"sites"`
	}
	a := &AutoGenerated{}
	err := gloson.Bind(s, a)
	if err == nil {
		fmt.Printf("%+v", a)
	}
}

//简单序列化
func Test3(t *testing.T) {
	type AutoGenerated struct {
		Name string `json:"name"`
		Num  int    `json:"num"`
	}
	a := &AutoGenerated{
		Name: "111",
		Num:  123,
	}
	s := gloson.Marshall(a)
	fmt.Println(s)
}

//简单float序列化
func Test4(t *testing.T) {
	type AutoGenerated struct {
		Name bool    `json:"name"`
		Num  float64 `json:"num"`
	}
	a := &AutoGenerated{
		Name: false,
		Num:  123.001,
	}
	s := gloson.Marshall(a)
	fmt.Println(s)
}

//简单数组序列化//带中文和中文字符的会爆炸
func Test5(t *testing.T) {
	type AutoGenerated struct {
		Name bool      `json:"name"`
		Num  []float64 `json:"num"`
		Test []string  `json:"test"`
	}
	a := &AutoGenerated{
		Name: false,
		Num:  []float64{123120.123, 4.2317823019, 6, 8273},
		Test: []string{"", "~!!@##$%%^&^&*(){}:}|", "asiojdoa"},
	}
	s := gloson.Marshall(a)
	fmt.Println(s)
}

//单层套娃
func TestUnmarshall_Bind(t *testing.T) {
	s := "{\n\"GlossDef\": {\n\"para\": \"A meta-markup language, used to create markup languages such as DocBook.\",\n\"GlossSeeAlso\": \"1\"\n},\n\"GlossSee\": \"markup\"\n}"
	type GlossDef struct {
		Para         string `json:"para"`
		GlossSeeAlso string `json:"GlossSeeAlso"`
	}
	type AutoGenerated struct {
		GlossDef GlossDef `json:"GlossDef"`
		GlossSee string   `json:"GlossSee"`
	}
	a := &AutoGenerated{}
	gloson.Bind(s, a)
	fmt.Println(a)

}

//对reflect包的测试
func TestReflect(t *testing.T) {
	//s := "{\n\"GlossDef\": {\n\"para\": \"A meta-markup language, used to create markup languages such as DocBook.\",\n\"GlossSeeAlso\": [\n\"GML\",\n\"XML\"\n]\n},\n\"GlossSee\": \"markup\"\n}"
	type GlossDef struct {
		Para         string   `json:"para"`
		GlossSeeAlso []string `json:"GlossSeeAlso"`
	}
	type AutoGenerated struct {
		GlossDef GlossDef `json:"GlossDef"`
		GlossSee string   `json:"GlossSee"`
	}
	a := &AutoGenerated{}
	va := reflect.ValueOf(a).Elem().Field(1)
	va.Set(reflect.ValueOf("1231231231"))
	fmt.Println("va", va)
	typ := reflect.TypeOf(a).Elem().FieldByIndex([]int{0, 1}).Tag.Get("json")
	fmt.Println("typ", typ)
	va2 := reflect.ValueOf(a).Elem().FieldByIndex([]int{1})
	va2.Set(reflect.ValueOf("wo"))
	fmt.Println(va2)
	//fmt.Println(reflect.TypeOf())

}

//三层结构体套娃 反序列化
func TestBind(t *testing.T) {
	s := "{\n\"GlossDef\": {\n\"para\": \"as DocBook.\",\n\"GlossSeeAlso\": {par:\"pppapp\",glo:1}\n},\n\"GlossSee\": \"markup\"\n}"
	type GlossSeeAlso struct {
		Par string `json:"par"`
		Glo int    `json:"glo"`
	}
	type GlossDef struct {
		Para         string       `json:"para"`
		GlossSeeAlso GlossSeeAlso `json:"GlossSeeAlso"`
	}
	type AutoGenerated struct {
		GlossDef GlossDef `json:"GlossDef"`
		GlossSee string   `json:"GlossSee"`
	}
	a := &AutoGenerated{}
	err := gloson.Bind(s, a)
	if err == nil {
		fmt.Printf("%+v", a)
	}
}

func TestMap(t *testing.T) {
	s := "{\n\"glossary\": {\n\"title\": \"example glossary\",\n\"GlossDiv\": {\n\"titles\": \"S\",\n\"GlossList\": {\n\"GlossEntry\": {\n\"ID\": \"SGML\",\n\"SortAs\": \"SGML\",\n\"GlossTerm\": \"Standard Generalized Markup Language\",\n\"Acronym\": \"SGML\",\n\"Abbrev\": \"ISO 8879:1986\",\n\"GlossDef\": {\n\"para\": \"A meta-markup language, used to create markup languages such as DocBook.\",\n\"GlossSeeAlso\": [\n\"GML\",\n\"XML\"\n]\n},\n\"GlossSee\": \"markup\"\n}\n}\n}\n}\n}"
	ma, _ := gloson.Map(s)
	fmt.Println(ma)
}

func TestMarshall2(t *testing.T) {
	type GlossDef struct {
		Para         string `json:"para"`
		GlossSeeAlso string `json:"GlossSeeAlso"`
		B            int    `json:"b"`
	}
	type AutoGenerated struct {
		GlossSee string   `json:"GlossSee"`
		GlossDef GlossDef `json:"GlossDef"`
		G        GlossDef `json:"g"`
		F        int      `json:"f"`
	}
	a := &AutoGenerated{
		GlossDef: GlossDef{
			Para:         "para",
			GlossSeeAlso: "ASDA",
			B:            9,
		},
		GlossSee: "3344",
		F:        123,
		G: GlossDef{
			Para:         "231",
			GlossSeeAlso: "223",
			B:            124,
		},
	}
	js := gloson.Marshall(a)
	fmt.Println(js)
}

func TestMarshall3(t *testing.T) {
	type GlossDiv struct {
		Title2    string `json:"title2"`
		GlossList int    `json:"GlossList"`
	}
	type Glossary struct {
		GGH      GlossDiv `json:"ggh"`
		Title1   string   `json:"title1"`
		K        string   `json:"k"`
		GlossDiv GlossDiv `json:"GlossDiv"`
	}
	type AutoGenerated struct {
		B        int      `json:"b"`
		Glossary Glossary `json:"Glossary"`
		C        string   `json:"c"`
		BDD      GlossDiv `json:"bdd"`
	}
	a := &AutoGenerated{
		Glossary: Glossary{
			GlossDiv: GlossDiv{
				Title2:    "23333",
				GlossList: 1,
			},
			Title1: "11415",
			K:      "adsasd",
			GGH: GlossDiv{
				Title2:    "23",
				GlossList: 90,
			},
		},
		B: 1,
		C: "123123",
		BDD: GlossDiv{
			Title2:    "231",
			GlossList: 1110,
		},
	}
	js := gloson.Marshall(a)
	fmt.Println(js)
}

func Benchmark(b *testing.B) {
	type GlossDiv struct {
		Title2    string `json:"title2"`
		GlossList int    `json:"GlossList"`
	}
	type Glossary struct {
		Title1   string   `json:"title1"`
		K        string   `json:"k"`
		GlossDiv GlossDiv `json:"GlossDiv"`
	}
	type AutoGenerated struct {
		B        int      `json:"b"`
		Glossary Glossary `json:"Glossary"`
		C        string   `json:"c"`
	}
	a := &AutoGenerated{
		Glossary: Glossary{
			GlossDiv: GlossDiv{
				Title2:    "23333",
				GlossList: 1,
			},
			Title1: "11415",
			K:      "adsasd",
		},
		B: 1,
		C: "123123",
	}
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		gloson.Marshall(a)
	}
}

func Benchmark2(b *testing.B) {
	var tc under.TypeChanger
	type GlossDiv struct {
		Title2    string `json:"title2"`
		GlossList int    `json:"GlossList"`
	}
	type Glossary struct {
		Title1   string   `json:"title1"`
		K        string   `json:"k"`
		GlossDiv GlossDiv `json:"GlossDiv"`
	}
	type AutoGenerated struct {
		B        int      `json:"b"`
		Glossary Glossary `json:"Glossary"`
		C        string   `json:"c"`
	}
	a := &AutoGenerated{
		Glossary: Glossary{
			GlossDiv: GlossDiv{
				Title2:    "23333",
				GlossList: 1,
			},
			Title1: "11415",
			K:      "adsasd",
		},
		B: 1,
		C: "123123",
	}
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		er, _ := json.Marshal(a)
		tc.Bytes2str(er)
	}
}

func Benchmark3(b *testing.B) {
	var tc under.TypeChanger
	s := "{\n\"GlossDef\": {\n\"para\": \"as DocBook.\",\n\"GlossSeeAlso\": {par:\"pppapp\",glo:1}\n},\n\"GlossSee\": \"markup\"\n}"
	type GlossSeeAlso struct {
		Par string `json:"par"`
		Glo int    `json:"glo"`
	}
	type GlossDef struct {
		Para         string       `json:"para"`
		GlossSeeAlso GlossSeeAlso `json:"GlossSeeAlso"`
	}
	type AutoGenerated struct {
		GlossDef GlossDef `json:"GlossDef"`
		GlossSee string   `json:"GlossSee"`
	}
	a := &AutoGenerated{}
	//for i := 0; i < b.N; i++ { // b.N，测试循环次数
	//	gloson.Bind(s, a)
	//}
	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		json.Unmarshal(tc.Str2bytes(s), a)
	}
}

func TestMarshall5(t *testing.T) {
	type GlossDiv struct {
		Title2    string `json:"title2"`
		GlossList int    `json:"GlossList"`
	}
	type Glossary struct {
		GGH      GlossDiv `json:"ggh"`
		Title1   string   `json:"title1"`
		K        string   `json:"k"`
		GlossDiv GlossDiv `json:"GlossDiv"`
	}
	type AutoGenerated struct {
		B        int      `json:"b"`
		Glossary Glossary `json:"Glossary"`
		C        string   `json:"c"`
		BDD      GlossDiv `json:"bdd"`
	}
	a := &AutoGenerated{
		Glossary: Glossary{
			GlossDiv: GlossDiv{
				Title2:    "23333",
				GlossList: 1,
			},
			Title1: "11415",
			K:      "adsasd",
			GGH: GlossDiv{
				Title2:    "23",
				GlossList: 90,
			},
		},
		B: 1,
		C: "123123",
		BDD: GlossDiv{
			Title2:    "231",
			GlossList: 1110,
		},
	}
	b := &AutoGenerated{}
	js := gloson.Marshall(a)
	gloson.Bind(js, b)
	fmt.Println(b)
}

func Test21(t *testing.T) {
	s := "{\"上海\":[\"浦东\",\"234\"],\"四川\":[\"成都\",\"攀枝花\"],\"福建\":[\"福州\",\"厦门\",\"泉州\"]}"
	type AutoGenerated struct {
		NamingFailed1 []string `json:"上海"`
		NamingFailed2 []string `json:"四川"`
		NamingFailed3 []string `json:"福建"`
	}
	a := &AutoGenerated{}
	err := gloson.Bind(s, a)
	if err == nil {
		fmt.Printf("%+v", a)
	}
}
