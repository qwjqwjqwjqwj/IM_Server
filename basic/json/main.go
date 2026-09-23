package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email,omitempty"`
}

func main() {
	//JSONstr 编码
	fmt.Println("EnCoding Json.....")
	u1 := User{123456, "qwj", "qwj12231", ""}
	fmt.Println(u1)
	JSONstr, err := json.Marshal(u1)
	if err != nil {
		fmt.Println("JOSN Marshal失败")
	}

	fmt.Println(JSONstr, "///type:", reflect.TypeOf(JSONstr))
	fmt.Printf("JSONstr=%s\n", JSONstr)

	//JSONstr 解码
	fmt.Println("DeCoding Json......")
	u2 := User{}
	err2 := json.Unmarshal(JSONstr, &u2)
	if err2 != nil {
		fmt.Println("DeCoding Json Failed")
	}
	fmt.Println(u2)
}
