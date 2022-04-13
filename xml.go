package main

import (
	"encoding/xml" //加载xml的库

	"fmt"

	"io/ioutil"

	"os"
)

type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"` //这里好像是固定字段必须这么写

	Version string `xml:"version.attr"`

	Svs []server `xml:"server"`

	Description string `xml:",innerxml"`
}

type server struct {
	XMLName xml.Name `xml:"server"`

	ServerName string `xml:"serverName"`

	ServerIP string `xml:"serverIP"`
}

func main() {

	//打开xml文件

	file, err := os.Open("test.xml")

	if err != nil {

		fmt.Printf("error:%v", err)

		return

	}

	defer file.Close()

	//读取文件

	data, err := ioutil.ReadAll(file)

	//fmt.Println(data)

	if err != nil {

		fmt.Printf("error:%v", err)

		return

	}

	v := Recurlyservers{}

	//解析成对应的struct对象

	err = xml.Unmarshal(data, &v)

	if err != nil {

		fmt.Printf("error: %v", err)

		return

	}

	fmt.Println(v)

}
