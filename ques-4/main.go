package main

import (
	"fmt"
	"encoding/json"
)

//https://stackoverflow.com/questions/48697961/unmarshal-2-different-structs-in-a-slice?answertab=votes#48701686

func main() {
	fmt.Println("Hello world")
	str := `[{
   "Url": "test.url",
   "Name": "testname"
},{
   "FormName": "Test - 2018",
   "FormNumber": 43,
   "FormSlug": "test-2018"
}]`
	type UrlData struct {
		Url  string `json:Url`
		Name string `json:Name`
	}

	type FormData struct {
		FormName   string `json:FormName`
		FormNumber int `json:FormNumber`
		FormSlug   string `json:FormSlug`
	}

	type ParallelData struct {
		UrlData
		FormData
	}

	var parallelData []ParallelData
	err := json.Unmarshal([]byte(str), &parallelData)
	if err != nil {
		panic(err)
	}
	fmt.Println(parallelData)
}
