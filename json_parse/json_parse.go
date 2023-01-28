package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 定义配置文件解析后的结构
type MongoConfig struct {
	MongoAddr       string
	MongoPoolLimit  int
	MongoDb         string
	MongoCollection string
}

type Config struct {
	Port  string
	Mongo MongoConfig
}

func main() {
	JsonParse := NewJsonStruct()
	v := Config{}
	JsonParse.Load("./json_parse.json", &v)
	fmt.Println(v.Port)
	fmt.Println(v.Mongo.MongoDb)
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}
func (js *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile()函数会读取文件的全部内容,并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为JSON格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
