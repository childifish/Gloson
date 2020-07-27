# GLOSON     :pray:希望人没事

🤣不知道是否高效但应该不高效的超简陋JSON框架(for golang)🙂

（dbq我不该觉得json框架简单）

## 🌳实现了

- ##### 简单JSON的反序列化:blush:

  ```go
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
  ```

  输出结果为

  ```
  === RUN   Test2
  &{Name:"网站" Num:3 Sites:[1 2 3]}--- PASS: Test2 (0.00s)
  PASS
  ```

- ##### 复杂JSON的反序列化（嵌套）:alien:

  ```go
  	s := "{\n\"GlossDef\": {\n\"para\": \"as DocBook.\",\n\"GlossSeeAlso\":	{par:\"pppapp\",glo:1}\n},\n\"GlossSee\": \"markup\"\n}"
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
  ```

  输出结果为

  ```
  === RUN   TestBind
  &{GlossDef:{Para:"asDocBook." GlossSeeAlso:{Par:"pppapp" Glo:1}} GlossSee:"markup"}--- PASS: TestBind (0.00s)
  PASS
  ```

- ##### 简单JSON的序列化:+1:

  ```go
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
  s, _ := gloson.Marshall(a)
  fmt.Println(s)
  ```

  输出结果为

  ```
  === RUN   Test5
  {
      "name":false,
      "num":[123120.123, 4.2317823019, 6, 8273],
      "test":["", "~!!@##$%%^&^&*(){}:}|", "asiojdoa"]
  }
  --- PASS: Test5 (0.00s)
  PASS
  ```

- ##### 比较复杂的JSON的序列化:jack_o_lantern:

  ```go
  type GlossDiv struct {
  	Title2    string `json:"title2"`
  	GlossList int    `json:"GlossList"`
  }
  type Glossary struct {
  	Title1   string   `json:"title1"`
  	GlossDiv GlossDiv `json:"GlossDiv"`
  	K        string   `json:"k"`
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
  js := gloson.Marshall(a)
  fmt.Println(js)
  ```

  输出结果为

  ```go
  === RUN   TestMarshall3
  {
      "b":1,
      "Glossary":{
          "title1":"11415",
          "GlossDiv":{
              "title2":"23333",
              "GlossList":1,
          },
          "k":"adsasd",
      },
      "c":"123123",
  }
  --- PASS: TestMarshall3 (0.00s)
  PASS
  ```

## :clap:实现过程

#### 序列化:v:

- 通过反射（reflect），可以让我们解析传入结构体，轻松（并不）地获取结构体的字段数（`.NumField()`），tag（`.Tag.Get("json")`），类型（`.Kind()`）

- 获取这些信息的前提是`reflect.ValueOf()`和`reflect.TypeOf()`其中`reflect.ValueOf()`必须以指针传递

- 获取精确的字段需要`.FieldByIndex()`或者`.Field()`，`.Field()`用int传参，`.FieldByIndex()`用int数组，两者的区别在于，`.FieldByIndex()`能够获取嵌套结构体内的成员，例如：

  ```go
  type GlossDiv struct {
  	Title2    string `json:"title2"`
  	GlossList int    `json:"GlossList"`
  }
  type Glossary struct {
  	Title1   string   `json:"title1"`
  	GlossDiv GlossDiv `json:"GlossDiv"`
  }
  type AutoGenerated struct {
  	Glossary Glossary `json:"Glossary"`
  	B        int      `json:"b"`
  }
  ```

  将AutoGenerated作为传入的参数，其中获取第一个成员Glossary需要的`.FieldByIndex([]int{0})`，而Glossary结构体内嵌在AutoGenerated中，Glossary的第一个成员Title1就是`.FieldByIndex([]int{0,0})`,GlossDiv里的GlossList则是`.FieldByIndex([]int{0,1,1})`,以此类推

- 获取了这些信息就可以自由组合了，序列化这个东西..因为是后写的，出的问题不多，而且确实简单，感觉有手就行吧，没什么好说的（加起来也就一百八十行）

- 序列化处理套娃的具体实现都写在代码里了

- 虽说能用但..

  这是自带的json序列化

  ```
  goos: windows
  goarch: amd64
  pkg: gloson/test
  Benchmark2
  Benchmark2-8      857743              1342 ns/op
  PASS
  ```

  这是我写的

  ```
  goos: windows
  goarch: amd64
  pkg: gloson/test
  Benchmark
  Benchmark-8        96774             14159 ns/op
  PASS
  ```

  大概有十倍的差距

- 而相对（绝对）困难的是下面的...

#### 反序列化​​😡

- 因为写之前担心太相似了，就没详细看json反序列化具体是怎么实现的，直接就开写了

- **实现逻辑**：传入两个参数，还未反序列化的json（string或者[]byte）和需要绑定的目标结构体

  1. 以逗号为标识，将字符串/[]byte分割为多个chuck

  2. 以冒号为标识，将每一个chuck分割为Key和value（冒号前和冒号后），并且根据value的类型（冒号后的第一个字符，比如数字或者负号开头一定是数字（int，float后续再细分），引号开头是字符串，”[“开头是数组，null,rue,fase这三个不在“”中，也可以识别）标识这个chuck，写入一个JsonKV结构体中，再写入[]JsonKV

     ```go
     type JsonKV struct {
     	Key        string
     	Within     interface{}
     	WithinType string
     }
     ```

  3. 然后解析[]JsonKV，用一个`CheckType()`函数判断之前写入的value（JsonKV.Within）,结合最初分析的类型确定这个value的具体类型，

     （其实这里已经可以开始着手绑定结构体了，但是我一开始脑子没转过来，另开了一个map[string]interface{}，反序列化的效率直接无了）

  4. 将jsonKV写入一张map，对未反序列化的json（string或者[]byte）的操作告一段落

  5. 然后开始对需要绑定的目标结构体进行操作，和序列化的操作类似，也是通过反射（reflect）来获取成员的位置，值和tag，因为涉及到结构体嵌套的问题，需要用到`.FieldByIndex()`，这里我用到了递归，在读取到的成员是结构体时，将该成员的位置作为参数传入`FindKeyD()`，这样就能解析到内层成员了

  6. 将每个字段的名字（因为涉及到成员名相同的问题，这里用tag获取），和位置，类型写入一张map

  7. 然后遍历两张map，当名字相同的时候，写入传入结构体中

  8. 聪明的小伙伴可能发现了，如果我先解析结构体再解析json分割后的每一个JsonKV struct，直接就可以`CheckType()`然后写入，只能说还是太年轻了

- 效率方面...这是我写的

  ```
  goos: windows
  goarch: amd64
  pkg: gloson/test
  Benchmark3
  Benchmark3-8       85110             23466 ns/op
  PASS
  ```

  这个是自带的

  ```
  goos: windows
  goarch: amd64
  pkg: gloson/test
  Benchmark3
  Benchmark3-8      923140              1549 ns/op
  PASS
  ```

  大概十五倍的差距

## 🐞还没完善的点

#### 序列化📃

- json库的序列化还没开始看
- 不支持成员为结构体数组的序列化
- map也不支持
- 常规的数组可以，但只有三种（string，int，float64）
- 写完了才发现正确的json每个对象的最后一对key-value都是没有逗号的
- GG

#### 反序列化📜

- json库看了一点点，大概知道是怎么实现反序列化的（用的一个decodeState结构体）
- return error非非非非非常不完善
- 不支持存在结构体数组的反序列化（前半部分json解析到map是可以的，后半部分解析结构体内的结构体数组还不行）
- 样本太少了，翻来覆去用我postman里留下的example里的json，都给薅秃了
- 测试tag为中文是否可行的时候发现数组里只有一个元素时有问题

## 🌈碎碎念time

- 嗯，写着写着就觉得自己这是面向脑淤血的编程，一个函数写下来有一万个变量v，有传入的`v interface{}`有遍历的`for k,v := range map{}`
- 没有了



#### <--to be continued...