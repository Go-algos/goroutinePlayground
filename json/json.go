package json

import (
	"encoding/json"
	"fmt"
)

// Refs: https://go.dev/blog/json

type Message struct {
	Name string
	Body string
	Time int64
}

func Runner() {
	m2 := Message{"Alice", "Hello", 1294706395881547000}
	jsonA, _ := json.Marshal(m2)
	fmt.Printf("Json: %s\n", jsonA)

	//json := []byte(`{"Name":"Bob","Food":"Pickle"}`)

	jsonAsByte := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(jsonAsByte, &f)
	fmt.Printf("Json2: %s\n", f)

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
