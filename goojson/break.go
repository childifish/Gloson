package goojson

var flag int

//拆分json

//将原始的byte数组以逗号区分，两个，之间的为一个item,带有{}的套娃视为一个item,返回string
func RawByte2Slice(rawBytes []byte)[]string  {
	markStart := 1
	//tag//这里用tag标识结构体套娃里的引号，当'{'和'}' ”抵消“时才分割//一开始用的++，--后来改成了bool取反   又改了回来---取反的话遇到"{...{"会爆炸
	//tag初始值为-1：因为最外层有”{“
	// TODO:要是字符串里面有{呢 解决：用”的个数判断
	// TODO:那要是“”中间有“呢 解决：用\判断---这个好像不用..不符合格式的string会报错(string里的”必须转义)
	tag := -1
	tag4tag := false
	var items []string
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
				items = append(items,bytes2str(rawBytes[markStart:i]))
				markStart = i+1
			}
		}
	}
	//len(rawBytes)-1可以去掉末尾的}
	items = append(items,bytes2str(rawBytes[markStart:len(rawBytes)-1]))
	//fmt.Println("item",items)
	return items
}

// todo:直接处理成完美形态
func BreakItem(items []string) []JsonRaw {
	flag ++
	//fmt.Println("flag",flag)
	//fmt.Println("break",items)
	var jsr JsonRaw
	var jsrSlice []JsonRaw
	for _,value := range items {
		itemB := str2bytes(value)
		for i,item := range itemB{
			//58是ASCII码的冒号：
			if item == 58 {
				jsr.Key = bytes2str(itemB[0:i])
				jsr.Within = bytes2str(itemB[i+1:])
				if (itemB[i+1]>=48&&itemB[i+1]<=57)||(itemB[i+1]==45) {
					//fmt.Println("number",bytes2str(itemB))
					jsr.WithinType = "number"
					break
				}
				switch itemB[i+1] {
				default:
					jsr.WithinType = "unknown"
					break
				case 123:// {
				//fmt.Println("option",bytes2str(itemB))
					jsr.WithinType = "option"
					break
				case 91:// [
					jsr.WithinType = "array"
					break
				case 34:// "
					jsr.WithinType = "string"
					break
				case 102,116,110:// f t n
					if i+5 <=len(itemB)-1{
						t := bytes2str(itemB[i+1:i+5])
						if  t == "true"{
							jsr.WithinType = "bool"
							break
						}
						n := bytes2str(itemB[i+1:i+5])
						if  n == "null"{
							jsr.WithinType = "null"
							break
						}
						f := bytes2str(itemB[i+1:i+6])
						if  f == "false"{
							jsr.WithinType = "bool"
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

