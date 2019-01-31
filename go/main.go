package main

import (
	"encoding/json"
	"fmt"
)

type test struct {
	A string `json:"a"`
	B int    `json:"b"`
}

func main() {
	cl, err := NewClient(&Option{Host: "127.0.0.1", Port: 9000})
	if err != nil {
		panic(err)
	}

	cl.Set("test", test{A: "this is value of field 'A'", B: 67})

	res := cl.Get("test")
	fmt.Println(string(res)) // will print => {"a":"this is value of field 'A'","b":67}

	var t test
	json.Unmarshal(res, &t)
	fmt.Printf("%+v\n", t)
}
