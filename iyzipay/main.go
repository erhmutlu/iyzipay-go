package main

import (
	"encoding/json"
)


type Server struct {
	Name    string `json:"name"`
}



func main() {
	s := &Server{
		Name:    "gopher",
	}

	var x map[string]interface{}
	inrec, _ := json.Marshal(s)
	json.Unmarshal(inrec, &x)

	m := make(map[string]string)
	m["name"] = "erhan"
	println(x["Name"])
}
