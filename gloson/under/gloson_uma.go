package under

type GlosonUma struct {
	RawByte GlosonByte	//传入的json bytes
	Buffer Buffer
	KvSlice []JsonKV	//
	Map map[string]interface{}
	BindBefore interface{}
	BindAfter interface{}
}

type Buffer struct {
	Items []string
}


func (glo *GlosonUma)StartMap(json string)error  {
	glo.RawByte = tc.Str2bytes(cl.CleanSpace(json))

	err := cl.CheckJson(glo.RawByte)
	if err!=nil{
		return err
	}

	glo.Glo2KV()

	err  = glo.WriteInMap()
	if err!=nil{
		return err
	}

	return nil
}

func (glo *GlosonUma)StartBinding(json string, v interface{})error{
	err := glo.StartMap(json)
	if err!=nil {
		return err
	}
	RecursionBinding(glo.Map,v)
	return nil
}



func (glo *GlosonUma) Glo2KV(){
	glo.Buffer.Items = glo.RawByte.Breaker()
	glo.KvSlice = glo.RawByte.Analysts(glo.Buffer.Items)
}

//将JsonKV全部写入FinalMap
func (glo *GlosonUma) WriteInMap()error{
	//fmt.Println("writer", jsrSlice)
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