package main

import (
	"encoding/json"
	"fmt"
)

type W struct {
	Username string `json:"username"`
	Age  int  `json:"age"`
	Sex  string

}

func main()  {
	user := &W{
		Username: "01",
		Sex: "nv",
		Age: 18,
	}
	data, _ := json.Marshal(user)
	fmt.Printf("json str:%s", string(data))
}





















