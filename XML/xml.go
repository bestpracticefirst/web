package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)
type User struct {
	Name string "user name"
	Passwd string "user password"
}
type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}
func main() {
	file, err := os.Open("XML/servers.xml") // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(v)
	user :=&User{"zhangsan","aaaaa"}
	s:=reflect.TypeOf(user).Elem()
	for i:=0;i<s.NumField();i++{
		fmt.Println(s.Field(i).Tag)
	}
}