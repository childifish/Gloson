package under

//反序列化
type GlosonUma struct {
	RawByte GlosonByte             //传入的json bytes
	Buffer  Buffer                 //切分后的Chuck
	KvSlice []JsonKV               //还未解析的JsonKV
	Map     map[string]interface{} //解析到Map
	BindingMsg
}

type Buffer struct {
	Items []string
}

//反序列化到map
func (glo *GlosonUma) StartMap(json string) error {
	//写入glo.rawByte, 清除json里的空格和换行符
	glo.RawByte = tc.Str2bytes(cl.CleanSpace(json))

	//判断是否是json
	//todo:多个json可以吗
	err := cl.CheckJson(glo.RawByte)
	if err != nil {
		return err
	}
	//开始分割
	glo.Glo2KV()

	err = glo.WriteInMap()
	if err != nil {
		return err
	}
	return nil
}

//开始绑定
func (glo *GlosonUma) StartBinding(json string, v interface{}) error {
	//先拆成map
	err := glo.StartMap(json)
	if err != nil {
		return err
	}
	map1 := glo.Map

	var b BindingMsg

	//初始化b的信息--写到map里
	//这里将传入结构体的key写入了surfaceKey
	b.InitBindingMsg(v)

	//遍历json的key-value组成的map
	for _, value := range map1 {
		//断言，这里应该不可能不是JsonKV，所以没有判断
		kv := value.(JsonKV)
		//json里的key都是带双引号的，这里要去掉
		key := cl.CleanMark(kv.Key)
		//通过json里的key和结构体的tag,找到这条value应该被写入的位置
		field, _ := Checkin(b.SurfaceKey, key)
		//最终的写入
		b.WriteItem(kv, field.Position)
	}
	return nil
}

//
func (glo *GlosonUma) Glo2KV() {
	//breaker将一个chuck拆分，
	glo.Buffer.Items = glo.RawByte.Breaker()
	glo.KvSlice = glo.RawByte.Analysts(glo.Buffer.Items)
}

//将JsonKV全部写入FinalMap
func (glo *GlosonUma) WriteInMap() error {
	kvSlice := glo.KvSlice
	var valueMap = make(map[string]interface{})
	for _, v := range kvSlice {
		key := v.Key
		err := v.CheckType()
		if err != nil {
			return err
		}
		valueMap[key] = v
	}
	glo.Map = valueMap
	return nil
}
