package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/api/core/v1"
	"io/ioutil"
	"path/filepath"
	"os"
)

func main()  {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	byt, err := ioutil.ReadFile(filepath.Join(wd, "ques-5/test.yaml"))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byt))

	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode(byt, nil, nil)
	c := obj.(*v1.List)

	fmt.Println(string(c.Items[0]))
}
