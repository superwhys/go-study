package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func judgeError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

/*
	go中对json进行序列化和反序列化的基本操作使用的是json.Marshal 和 json.Unmarshal
 */

type Person struct {
	Name string
	Age int
	Weight float64
}

func BaseJsonOperation() {
	p1 := Person{
		"Why",
		18,
		71.5,
	}

	// 利用json.Marshal 序列化
	Json2p, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("str: %s\n", Json2p)

	// 利用json.Unmarshal对pJson进行反序列化
	var p2Json Person

	err = json.Unmarshal(Json2p, &p2Json)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(p2Json)
}

// Person2 利用tag给某个字段打上标签，序列化后的字段名就是tag中的名字， omitempty是当字段为空是忽略该字段
// `json:"-"` 表示序列化时忽略这个字段
type Person2 struct {
	Name   string  `json:"MyName,omitempty"`
	Age    int	   `json:"MyAge,omitempty"`
	Weight float64 `json:"-"` // 表示忽略该字段

}

func TagInJson() {
	p1 := Person2 {
		"why",
		18,
		71.4,
	}

	Json2Str, err := json.Marshal(p1)
	judgeError(err)
	fmt.Println(string(Json2Str))
}

// House 忽略嵌套结构体中的空值
type House struct {
	Address string
	Owner Person2
	// 当嵌套的结构体中不使用tag时，序列化后的json是单层的，若要序列化后也是嵌套的格式，需要添加tag或使用具名嵌套
	Person2
	Person `json:"Owner2"`
}

func NestedStruct() {
	h1 := House {
		"GuangDong Shenzhen",
		Person2 {
			"杨浩文",
			18,
			75.1,
		},
		Person2 {
			"why",
			18,
			71.4,
		},
		Person{
			Name:   "yhx",
			Age:    20,
			Weight: 60,
		},
	}

	Json2Str, err := json.Marshal(h1)
	judgeError(err)
	fmt.Printf("%s\n", Json2Str)
	// {"Address":"GuangDong Shenzhen","Owner":{"MyName":"杨浩文","MyAge":18},"MyName":"why","MyAge":18}
}

type Phone struct {
	PhoneNum string `json:"phone-num,omitempty"`
	Owner *Person2  `json:"owner,omitempty"`
}

// IgnoreNestedStruct 想要在嵌套的结构体为空值时，忽略该字段，
// 仅添加omitempty是不够的, 还需要使用嵌套的结构体指针
func IgnoreNestedStruct() {
	phone := Phone {
		PhoneNum: "1234567890",
	}
	phone1 := Phone {
		PhoneNum: "1234567890",
		Owner: &Person2 {
			Name: "yhw",
			Age: 18,
		},
	}
	Json2Str, err := json.Marshal(phone)
	judgeError(err)
	fmt.Printf("%s\n", Json2Str)

	Json2Str1, err := json.Marshal(phone1)
	judgeError(err)
	fmt.Printf("%s\n", Json2Str1)
}

type User struct {
	Name 	 string `json:"name"`
	Password string `json:"password"`
}

// PubUser 我们需要json序列化User，但是不想把密码也序列化，
// 又不想修改User结构体，这个时候我们就可以使用创建另外一个结构体PublicUser匿名嵌套原User，
// 同时指定Password字段为匿名结构体指针类型，并添加omitempty tag
type PubUser struct {
	*User // 必须是匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}

func IgnoreNullValueWithNoModify() {
	u1 := User {
		Name: "why",
		Password: "123456",
	}

	Json2Str, err := json.Marshal(PubUser{User: &u1})
	judgeError(err)
	fmt.Printf("%s\n", Json2Str)
}

// Card 有时候，前端在传递来的json数据中可能会使用字符串类型的数字，
// 这个时候可以在结构体tag中添加string来告诉json包从字符串中解析相应字段的数据
type Card struct {
	Id    int     `json:"id,string"`
	Score float64 `json:"score,string"`
}

func TypeTag() {
	jsonStr1 := `{"id": "1234567","score": "88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("c1:%#v\n", c1) // c1:main.Card{ID:1234567, Score:88.5}
}

func decoderDemo() {
	// map[string]interface{} -> json string
	var m = make(map[string]interface{}, 1)
	m["count"] = 1 // int
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
	}
	fmt.Printf("str:%#v\n", string(b))
	// json string -> map[string]interface{}
	var m2 map[string]interface{}
	// 使用decoder方式反序列化，指定使用number类型
	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.UseNumber()
	err = decoder.Decode(&m2)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("value:%v\n", m2["count"]) // 1
	fmt.Printf("type:%T\n", m2["count"])  // json.Number
	// 将m2["count"]转为json.Number之后调用Int64()方法获得int64类型的值
	count, err := m2["count"].(json.Number).Int64()
	if err != nil {
		fmt.Printf("parse to int64 failed, err:%v\n", err)
		return
	}
	fmt.Printf("type:%T, %v\n", count, count) // int
}

type sendMsg struct {
	User string `json:"user"`
	Msg  string `json:"msg"`
}

func rawMessageDemo() {
	jsonStr := `{"sendMsg":{"user":"q1mi","msg":"永远不要高估自己"},"say":"Hello"}`
	// 定义一个map，value类型为json.RawMessage，方便后续更灵活地处理
	var data map[string]json.RawMessage
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Printf("json.Unmarshal jsonStr failed, err:%v\n", err)
		return
	}
	fmt.Printf("%v\n", data)
	var msg sendMsg
	if err := json.Unmarshal(data["sendMsg"], &msg); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("msg:%#v\n", msg)
	// msg:main.sendMsg{User:"q1mi", Msg:"永远不要高估自己"}
}

func main() {
	BaseJsonOperation()
	TagInJson()
	NestedStruct()
	IgnoreNestedStruct()
	IgnoreNullValueWithNoModify()
	TypeTag()
	decoderDemo()
	rawMessageDemo()
}
