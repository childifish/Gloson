package under

import "fmt"

type GoojsonByte []byte

//将goo分解成chuck，再解析chuck-item到结构体数组
func (g *GoojsonByte)Go2KV()[]JsonKV  {
	items :=  g.Breaker()
	return g.Analysts(items)
}

func (g *GoojsonByte)Go2Array()[][]JsonKV  {
	items :=  g.EXBreaker()
	var ite GoojsonByte
	var jsonkv []JsonKV
	var jsonkvs [][]JsonKV
	for _,v := range items{
		ite = tc.Str2bytes(v)
		fmt.Print(v)
		jsonkv = ite.Analysts(ite.Breaker())
		jsonkvs = append(jsonkvs,jsonkv)
	}
	return jsonkvs
}

//返回goojson本身
func (g *GoojsonByte)Self()[]byte  {
	return *g
}

//破坏者 将总的goojson[]byte拆分为chuck
//chuck是json中的一组还未解析的Key-Value
func (g *GoojsonByte)Breaker()[]string  {
	var items []string
	rawBytes := g.Self()
	markStart := 1
	//tag//这里用tag标识结构体套娃里的引号，当'{'和'}' ”抵消“时才分割//一开始用的++，--后来改成了bool取反   又改了回来---取反的话遇到"{...{"会爆炸
	//tag初始值为-1：因为最外层有”{“
	// 要是字符串里面有{呢 解决：用”的个数判断
	// 那要是“”中间有“呢 解决：用\判断---这个好像不用..不符合格式的string会报错(string里的”必须转义)
	// 如果是数组呢
	tag := -1
	tag4tag := false

	for i,byt := range rawBytes{
		//if byt
		//34是ASCII码的引号"
		if byt == 34{
			//不用&&大多数时候能少一次判断
			if rawBytes[i-1] == 92{//92是ASCII码的反斜杠\
				//当 前一位byte为转义符时，此次"不计入
				continue
			}
			//取反
			tag4tag = !tag4tag
			continue
		}
		//当tag4tag为真--即在双引号里，跳过本次循环（可以防止字符串和key里的特殊符号）
		if tag4tag {
			continue
		}
		switch byt{
		case 123:
			tag ++
			continue
		case 125:
			tag --
			continue
		case 44://44是ASCII码的逗号,
			if tag ==0{
				items = append(items,tc.Bytes2str(rawBytes[markStart:i]))
				markStart = i+1
			}
		}
	}
	//len(rawBytes)-1可以去掉末尾的}
	items = append(items,tc.Bytes2str(rawBytes[markStart:len(rawBytes)-1]))
	//fmt.Println("item",items)
	return items
}

//用来拆分数组
func (g *GoojsonByte)EXBreaker()[]string  {
	fmt.Println("-------------------------exbreaker----------------")
	var items []string
	rawBytes := g.Self()
	fmt.Println(tc.Bytes2str(rawBytes))
	markStart := 1
	//tag//这里用tag标识结构体套娃里的引号，当'{'和'}' ”抵消“时才分割//一开始用的++，--后来改成了bool取反   又改了回来---取反的话遇到"{...{"会爆炸
	//tag初始值为-1：因为最外层有”{“
	// 要是字符串里面有{呢 解决：用”的个数判断
	// 那要是“”中间有“呢 解决：用\判断---这个好像不用..不符合格式的string会报错(string里的”必须转义)
	// 如果是数组呢
	tag := 0
	tag4tag := false

	for i,byt := range rawBytes{
		//if byt
		//34是ASCII码的引号"
		if byt == 34{
			//不用&&大多数时候能少一次判断
			if i>1{
				if rawBytes[i-1] == 92{//92是ASCII码的反斜杠\
					//当 前一位byte为转义符时，此次"不计入
					continue
				}
				//取反
				tag4tag = !tag4tag
				continue
			}
			continue
		}
		//当tag4tag为真--即在双引号里，跳过本次循环（可以防止字符串和key里的特殊符号）
		if tag4tag {
			continue
		}
		switch byt{
		case 123:
			tag ++
			continue
		case 125:
			tag --
			continue
		case 44://44是ASCII码的逗号,
			if tag ==0{
				items = append(items,tc.Bytes2str(rawBytes[markStart:i+1]))
				markStart = i+1
				for k,v := range items {
					fmt.Println("inner",k,v)
				}
			}
		}
	}
	//len(rawBytes)-1可以去掉末尾的}
	items = append(items,tc.Bytes2str(rawBytes[markStart:len(rawBytes)-1]))
	//fmt.Println("item",items)
	return items
}

//分析器 解析value
//以：分割K和V，再根据V开头的字节判断V的类型
func (g *GoojsonByte)Analysts(items []string) []JsonKV {
	var jsr JsonKV
	var jsrSlice []JsonKV
	for _,value := range items {
		itemB := tc.Str2bytes(value)
		for i,item := range itemB{
			//58是ASCII码的冒号：
			if item == 58 {
				jsr.Key = tc.Bytes2str(itemB[0:i])
				jsr.Within = tc.Bytes2str(itemB[i+1:])
				//负号以及数字开头
				if ((itemB[i+1]>=48)&&(itemB[i+1]<=57))||(itemB[i+1]==45) {
					jsr.WithinType = "number"
					break
				}
				switch itemB[i+1] {
				default:
					jsr.WithinType = "unknown"
					break
				case 123:// {
					jsr.WithinType = "option"
					break
				case 91:// [
					jsr.WithinType = "array"
					break
				case 34:// "
					jsr.WithinType = "string"
					break
				case 102,116,110:// f t n
					if i+5 <=len(itemB){
						t := tc.Bytes2str(itemB[i+1:i+5])
						if  t == "true"{
							jsr.WithinType = "true"
							break
						}
						n := tc.Bytes2str(itemB[i+1:i+5])
						if  n == "null"{
							jsr.WithinType = "null"
							break
						}
						f := tc.Bytes2str(itemB[i+1:i+6])
						if  f == "false"{
							jsr.WithinType = "false"
							break
						}
						jsr.WithinType = "unknown"
					}
					break
				}
				break
			}
		}
		jsrSlice = append(jsrSlice,jsr)
	}
	return jsrSlice
}
