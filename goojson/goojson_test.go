package goojson

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringHello(t *testing.T)  {
	s := "hello"
	print(s)
}

func TestBytes(t *testing.T)  {
	s := "{hello}"
	rawBytes := str2bytes(s)
	for i,j := range rawBytes{
		fmt.Println(i,j)
	}
	fmt.Println(len(rawBytes))
}

func TestBytes1(t *testing.T)  {
	s := "hello}"
	rawBytes := str2bytes(s)
	if rawBytes[0]!=123||rawBytes[len(rawBytes)-1]!=125{
		fmt.Println("err json")
	}else {
		fmt.Println("json")
	}
	s2 := "{hello}"
	rawBytes2 := str2bytes(s2)
	if rawBytes2[0]!=123||rawBytes2[len(rawBytes2)-1]!=125{
		fmt.Println("err json")
	}else {
		fmt.Println("json")
	}
}

func TestChinese(t *testing.T)  {
	s := "你好鸭！！！"
	rawBytes := str2bytes(s)
	for i,j := range rawBytes{
		fmt.Println(i,j)
	}
	fmt.Println(len(rawBytes))
	fmt.Println(rawBytes)
	notraw := bytes2str(rawBytes)
	fmt.Println(notraw)
}

func TestChineseCheck(t *testing.T)  {
	s := "{你好鸭！！！}"
	rawBytes := str2bytes(s)
	fmt.Println(len(rawBytes))
	if rawBytes[0]!=123||rawBytes[len(rawBytes)-1]!=125{
		fmt.Println("err json")
	}else {
		fmt.Println("json")
	}
	s2 := "你好鸭！！！"
	rawBytes2 := str2bytes(s2)
	if rawBytes2[0]!=123||rawBytes2[len(rawBytes2)-1]!=125{
		fmt.Println("err json")
	}else {
		fmt.Println("json")
	}
}

func TestRB2S(t *testing.T)  {
	s := "{\n    \"data\": {\n        \"goods\": null\n    },\n    \"info\": \"success\",\n    \"status\": 200\n}"
	re:=strings.ReplaceAll(s,"\n","")
	rawBytes := str2bytes(re)
	fmt.Println(rawBytes)
	if rawBytes[0]!=123||rawBytes[len(rawBytes)-1]==125{
		markStart := 1
		var itemB []byte
		var items []string
		for i,byt := range rawBytes{
			if byt == 44 {
				itemB = rawBytes[markStart:i]
				markStart = i
				items = append(items,bytes2str(itemB))
			}
		}
		for k,v := range items{
			fmt.Println(k,v)
		}
	}
}

func TestRB2S2(t *testing.T)  {
	s := "{\n    \"info\": \"success\",\n    \"item\": {\n        \"1\": {\n            \"ID\": 1,\n            \"Name\": \"EHLVEEXN\",\n            \"Total\": 192,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"2\": {\n            \"ID\": 2,\n            \"Name\": \"CGSRGMVC\",\n            \"Total\": 253,\n            \"Left\": 147,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"3\": {\n            \"ID\": 3,\n            \"Name\": \"IXNBFIET\",\n            \"Total\": 275,\n            \"Left\": 145,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"4\": {\n            \"ID\": 4,\n            \"Name\": \"CXSWQGAL\",\n            \"Total\": 346,\n            \"Left\": 226,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"5\": {\n            \"ID\": 5,\n            \"Name\": \"CWHEGWDZ\",\n            \"Total\": 50,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"6\": {\n            \"ID\": 6,\n            \"Name\": \"SDFUNJFG\",\n            \"Total\": 168,\n            \"Left\": 52,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"7\": {\n            \"ID\": 7,\n            \"Name\": \"CNKLUXFR\",\n            \"Total\": 461,\n            \"Left\": 339,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"8\": {\n            \"ID\": 8,\n            \"Name\": \"FZQDICTN\",\n            \"Total\": 390,\n            \"Left\": 268,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"9\": {\n            \"ID\": 9,\n            \"Name\": \"WZWROCQQ\",\n            \"Total\": 298,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"10\": {\n            \"ID\": 10,\n            \"Name\": \"OJKEAGFX\",\n            \"Total\": 138,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"11\": {\n            \"ID\": 11,\n            \"Name\": \"UPEVWGDN\",\n            \"Total\": 307,\n            \"Left\": 194,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"12\": {\n            \"ID\": 12,\n            \"Name\": \"XQIKIVDB\",\n            \"Total\": 65,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"13\": {\n            \"ID\": 13,\n            \"Name\": \"DKNBZTBL\",\n            \"Total\": 148,\n            \"Left\": 36,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"14\": {\n            \"ID\": 14,\n            \"Name\": \"VUCNVUIO\",\n            \"Total\": 156,\n            \"Left\": 84,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"15\": {\n            \"ID\": 15,\n            \"Name\": \"ONNHPSLS\",\n            \"Total\": 69,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"16\": {\n            \"ID\": 16,\n            \"Name\": \"QBQKAUGX\",\n            \"Total\": 161,\n            \"Left\": 69,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"17\": {\n            \"ID\": 17,\n            \"Name\": \"EZKMKDIS\",\n            \"Total\": 297,\n            \"Left\": 175,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"18\": {\n            \"ID\": 18,\n            \"Name\": \"KSOKSBZC\",\n            \"Total\": 57,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"19\": {\n            \"ID\": 19,\n            \"Name\": \"YBUOXUPN\",\n            \"Total\": 200,\n            \"Left\": 62,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"20\": {\n            \"ID\": 20,\n            \"Name\": \"DMMPAORN\",\n            \"Total\": 288,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"21\": {\n            \"ID\": 21,\n            \"Name\": \"SZBYCNYW\",\n            \"Total\": 124,\n            \"Left\": 124,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"22\": {\n            \"ID\": 22,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"23\": {\n            \"ID\": 23,\n            \"Name\": \"tttttedt22\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"24\": {\n            \"ID\": 24,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"25\": {\n            \"ID\": 25,\n            \"Name\": \"tedt1\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"26\": {\n            \"ID\": 26,\n            \"Name\": \"\",\n            \"Total\": 0,\n            \"Left\": 0,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        }\n    },\n    \"status\": 200\n}"
	rawBytes := str2bytes(s)
	if rawBytes[0]!=123||rawBytes[len(rawBytes)-1]!=125{
		fmt.Println("bad")
	}else {
		//有strings.Split()但还是想写个自己的
		//后来发现还真得自己写，不然不能处理结构体套娃
		items := RawByte2Slice(rawBytes)
		for k,v := range items{
			fmt.Println(k,v)
		}
	}
}

func TestRB2S3(t *testing.T)  {
	s := "{\n    \"data\": {\n        \"goods\": null\n    },\n    \"info\": \"success\",\n    \"status\": 200\n}"
	rawBytes := str2bytes(s)
	if rawBytes[0]!=123||rawBytes[len(rawBytes)-1]!=125{
		fmt.Println("bad")
	}else {
		//有strings.Split()但还是想写个自己的
		//后来发现还真得自己写，不然不能处理结构体套娃
		items := RawByte2Slice(rawBytes)
		for k,v := range items{
			fmt.Println(k,v)
		}
	}
}

func TestRB2S4(t *testing.T)  {
	s := "{\n    \"info\": \"success\",\n    \"item\": {\n        \"1\": {\n            \"ID\": 1,\n            \"Name\": \"EHLVEEXN\",\n            \"Total\": 192,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"2\": {\n            \"ID\": 2,\n            \"Name\": \"CGSRGMVC\",\n            \"Total\": 253,\n            \"Left\": 147,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"3\": {\n            \"ID\": 3,\n            \"Name\": \"IXNBFIET\",\n            \"Total\": 275,\n            \"Left\": 145,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"4\": {\n            \"ID\": 4,\n            \"Name\": \"CXSWQGAL\",\n            \"Total\": 346,\n            \"Left\": 226,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"5\": {\n            \"ID\": 5,\n            \"Name\": \"CWHEGWDZ\",\n            \"Total\": 50,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"6\": {\n            \"ID\": 6,\n            \"Name\": \"SDFUNJFG\",\n            \"Total\": 168,\n            \"Left\": 52,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"7\": {\n            \"ID\": 7,\n            \"Name\": \"CNKLUXFR\",\n            \"Total\": 461,\n            \"Left\": 339,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"8\": {\n            \"ID\": 8,\n            \"Name\": \"FZQDICTN\",\n            \"Total\": 390,\n            \"Left\": 268,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"9\": {\n            \"ID\": 9,\n            \"Name\": \"WZWROCQQ\",\n            \"Total\": 298,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"10\": {\n            \"ID\": 10,\n            \"Name\": \"OJKEAGFX\",\n            \"Total\": 138,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"11\": {\n            \"ID\": 11,\n            \"Name\": \"UPEVWGDN\",\n            \"Total\": 307,\n            \"Left\": 194,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"12\": {\n            \"ID\": 12,\n            \"Name\": \"XQIKIVDB\",\n            \"Total\": 65,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"13\": {\n            \"ID\": 13,\n            \"Name\": \"DKNBZTBL\",\n            \"Total\": 148,\n            \"Left\": 36,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"14\": {\n            \"ID\": 14,\n            \"Name\": \"VUCNVUIO\",\n            \"Total\": 156,\n            \"Left\": 84,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"15\": {\n            \"ID\": 15,\n            \"Name\": \"ONNHPSLS\",\n            \"Total\": 69,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"16\": {\n            \"ID\": 16,\n            \"Name\": \"QBQKAUGX\",\n            \"Total\": 161,\n            \"Left\": 69,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"17\": {\n            \"ID\": 17,\n            \"Name\": \"EZKMKDIS\",\n            \"Total\": 297,\n            \"Left\": 175,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"18\": {\n            \"ID\": 18,\n            \"Name\": \"KSOKSBZC\",\n            \"Total\": 57,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"19\": {\n            \"ID\": 19,\n            \"Name\": \"YBUOXUPN\",\n            \"Total\": 200,\n            \"Left\": 62,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"20\": {\n            \"ID\": 20,\n            \"Name\": \"DMMPAORN\",\n            \"Total\": 288,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"21\": {\n            \"ID\": 21,\n            \"Name\": \"SZBYCNYW\",\n            \"Total\": 124,\n            \"Left\": 124,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"22\": {\n            \"ID\": 22,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"23\": {\n            \"ID\": 23,\n            \"Name\": \"tttttedt22\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"24\": {\n            \"ID\": 24,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"25\": {\n            \"ID\": 25,\n            \"Name\": \"tedt1\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"26\": {\n            \"ID\": 26,\n            \"Name\": \"\",\n            \"Total\": 0,\n            \"Left\": 0,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        }\n    },\n    \"status\": 200\n}"
	rawBytes,_:= WhetherString2Bytes(s)
	if rawBytes[0]!=123||rawBytes[len(rawBytes)-1]!=125{
		fmt.Println("bad")
	}else {
		//有strings.Split()但还是想写个自己的
		//后来发现还真得自己写，不然不能处理结构体套娃
		items := RawByte2Slice(rawBytes)
		for k,v := range items {
			fmt.Println(k,v)
		}
		jsrSlice:= BreakItem(items)
		for k,v := range jsrSlice {
			fmt.Println(k,v)
		}
	}
}

func TestRB2S5(t *testing.T)  {
	s := "{\n    \"info\": \"success\",\n    \"item\": {\n        \"1\": {\n            \"ID\": 1,\n            \"Name\": \"EHLVEEXN\",\n            \"Total\": 192,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"2\": {\n            \"ID\": 2,\n            \"Name\": \"CGSRGMVC\",\n            \"Total\": 253,\n            \"Left\": 147,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"3\": {\n            \"ID\": 3,\n            \"Name\": \"IXNBFIET\",\n            \"Total\": 275,\n            \"Left\": 145,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"4\": {\n            \"ID\": 4,\n            \"Name\": \"CXSWQGAL\",\n            \"Total\": 346,\n            \"Left\": 226,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"5\": {\n            \"ID\": 5,\n            \"Name\": \"CWHEGWDZ\",\n            \"Total\": 50,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"6\": {\n            \"ID\": 6,\n            \"Name\": \"SDFUNJFG\",\n            \"Total\": 168,\n            \"Left\": 52,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"7\": {\n            \"ID\": 7,\n            \"Name\": \"CNKLUXFR\",\n            \"Total\": 461,\n            \"Left\": 339,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"8\": {\n            \"ID\": 8,\n            \"Name\": \"FZQDICTN\",\n            \"Total\": 390,\n            \"Left\": 268,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"9\": {\n            \"ID\": 9,\n            \"Name\": \"WZWROCQQ\",\n            \"Total\": 298,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"10\": {\n            \"ID\": 10,\n            \"Name\": \"OJKEAGFX\",\n            \"Total\": 138,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"11\": {\n            \"ID\": 11,\n            \"Name\": \"UPEVWGDN\",\n            \"Total\": 307,\n            \"Left\": 194,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"12\": {\n            \"ID\": 12,\n            \"Name\": \"XQIKIVDB\",\n            \"Total\": 65,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"13\": {\n            \"ID\": 13,\n            \"Name\": \"DKNBZTBL\",\n            \"Total\": 148,\n            \"Left\": 36,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"14\": {\n            \"ID\": 14,\n            \"Name\": \"VUCNVUIO\",\n            \"Total\": 156,\n            \"Left\": 84,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"15\": {\n            \"ID\": 15,\n            \"Name\": \"ONNHPSLS\",\n            \"Total\": 69,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"16\": {\n            \"ID\": 16,\n            \"Name\": \"QBQKAUGX\",\n            \"Total\": 161,\n            \"Left\": 69,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"17\": {\n            \"ID\": 17,\n            \"Name\": \"EZKMKDIS\",\n            \"Total\": 297,\n            \"Left\": 175,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"18\": {\n            \"ID\": 18,\n            \"Name\": \"KSOKSBZC\",\n            \"Total\": 57,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"19\": {\n            \"ID\": 19,\n            \"Name\": \"YBUOXUPN\",\n            \"Total\": 200,\n            \"Left\": 62,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"20\": {\n            \"ID\": 20,\n            \"Name\": \"DMMPAORN\",\n            \"Total\": 288,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"21\": {\n            \"ID\": 21,\n            \"Name\": \"SZBYCNYW\",\n            \"Total\": 124,\n            \"Left\": 124,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"22\": {\n            \"ID\": 22,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"23\": {\n            \"ID\": 23,\n            \"Name\": \"tttttedt22\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"24\": {\n            \"ID\": 24,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"25\": {\n            \"ID\": 25,\n            \"Name\": \"tedt1\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"26\": {\n            \"ID\": 26,\n            \"Name\": \"\",\n            \"Total\": 0,\n            \"Left\": 0,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        }\n    },\n    \"status\": 200\n}"
	//s := "{\n    \"data\": {\n        \"goods\": 123\n    },\n    \"info\": \"success\",\n    \"status\": 12345156\n}"

	rawBytes,_:= WhetherString2Bytes(s)
	if rawBytes[0]!=123||rawBytes[len(rawBytes)-1]!=125{
		fmt.Println("bad")
	}else {
		items := RawByte2Slice(rawBytes)
		jsrSlice:= BreakItem(items)
		fmt.Println(jsrSlice)
		kmap,_ := JsonRaw2KeyMap(jsrSlice)
		fmt.Println(kmap)
		for _,v := range kmap{
			fmt.Println(v)
		}
	}
}

func TestUnmarshall(t *testing.T)  {
	s := "{\n    \"info\": \"success\",\n    \"item\": {\n        \"1\": {\n            \"ID\": 1,\n            \"Name\": \"EHLVEEXN\",\n            \"Total\": 192,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"2\": {\n            \"ID\": 2,\n            \"Name\": \"CGSRGMVC\",\n            \"Total\": 253,\n            \"Left\": 147,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"3\": {\n            \"ID\": 3,\n            \"Name\": \"IXNBFIET\",\n            \"Total\": 275,\n            \"Left\": 145,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"4\": {\n            \"ID\": 4,\n            \"Name\": \"CXSWQGAL\",\n            \"Total\": 346,\n            \"Left\": 226,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"5\": {\n            \"ID\": 5,\n            \"Name\": \"CWHEGWDZ\",\n            \"Total\": 50,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"6\": {\n            \"ID\": 6,\n            \"Name\": \"SDFUNJFG\",\n            \"Total\": 168,\n            \"Left\": 52,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"7\": {\n            \"ID\": 7,\n            \"Name\": \"CNKLUXFR\",\n            \"Total\": 461,\n            \"Left\": 339,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"8\": {\n            \"ID\": 8,\n            \"Name\": \"FZQDICTN\",\n            \"Total\": 390,\n            \"Left\": 268,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"9\": {\n            \"ID\": 9,\n            \"Name\": \"WZWROCQQ\",\n            \"Total\": 298,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"10\": {\n            \"ID\": 10,\n            \"Name\": \"OJKEAGFX\",\n            \"Total\": 138,\n            \"Left\": 56,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"11\": {\n            \"ID\": 11,\n            \"Name\": \"UPEVWGDN\",\n            \"Total\": 307,\n            \"Left\": 194,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"12\": {\n            \"ID\": 12,\n            \"Name\": \"XQIKIVDB\",\n            \"Total\": 65,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"13\": {\n            \"ID\": 13,\n            \"Name\": \"DKNBZTBL\",\n            \"Total\": 148,\n            \"Left\": 36,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"14\": {\n            \"ID\": 14,\n            \"Name\": \"VUCNVUIO\",\n            \"Total\": 156,\n            \"Left\": 84,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"15\": {\n            \"ID\": 15,\n            \"Name\": \"ONNHPSLS\",\n            \"Total\": 69,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"16\": {\n            \"ID\": 16,\n            \"Name\": \"QBQKAUGX\",\n            \"Total\": 161,\n            \"Left\": 69,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"17\": {\n            \"ID\": 17,\n            \"Name\": \"EZKMKDIS\",\n            \"Total\": 297,\n            \"Left\": 175,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"18\": {\n            \"ID\": 18,\n            \"Name\": \"KSOKSBZC\",\n            \"Total\": 57,\n            \"Left\": 0,\n            \"IsSoldOut\": true,\n            \"Lock\": {}\n        },\n        \"19\": {\n            \"ID\": 19,\n            \"Name\": \"YBUOXUPN\",\n            \"Total\": 200,\n            \"Left\": 62,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"20\": {\n            \"ID\": 20,\n            \"Name\": \"DMMPAORN\",\n            \"Total\": 288,\n            \"Left\": 196,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"21\": {\n            \"ID\": 21,\n            \"Name\": \"SZBYCNYW\",\n            \"Total\": 124,\n            \"Left\": 124,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"22\": {\n            \"ID\": 22,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"23\": {\n            \"ID\": 23,\n            \"Name\": \"tttttedt22\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"24\": {\n            \"ID\": 24,\n            \"Name\": \"tttttedt\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"25\": {\n            \"ID\": 25,\n            \"Name\": \"tedt1\",\n            \"Total\": 998,\n            \"Left\": 998,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        },\n        \"26\": {\n            \"ID\": 26,\n            \"Name\": \"\",\n            \"Total\": 0,\n            \"Left\": 0,\n            \"IsSoldOut\": false,\n            \"Lock\": {}\n        }\n    },\n    \"status\": 200\n}"
	//s := "{\n    \"data\": {\n        \"goods\": null\n    },\n    \"info\": \"success\",\n    \"status\": 200\n}"
	jsr := GooJson{
		Json:    s,
		NotJson: nil,
		Err:     nil,
	}
	r,err := jsr.Unmarshall(s)
	result := r.NotJson
	if err!=nil {
		fmt.Println("result",result)
	}
}