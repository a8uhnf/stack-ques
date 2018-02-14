package main

import (
	"github.com/go-openapi/loads"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/analysis"
	"github.com/go-openapi/validate"
	"k8s.io/api/extensions/v1beta1"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
)

//https://stackoverflow.com/questions/48697961/unmarshal-2-different-structs-in-a-slice?answertab=votes#48701686

func main() {
	path := "/home/tigerworks/.kubepack/v1.9.3/openapi-spec/swagger.json"
	schema := validate.SchemaValidator{  // define a validate.SchemaValidator
		Path: path,                      // which contains your swagger.json file and schema
		Schema: spec.MustLoadSwagger20Schema(),
	}

	byt := []byte(`
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  labels:
    name: redis-master
  name: name
spec:
  replicas: "1"
  template:
    metadata:
      labels:
        app: redis
    spec:
      containers:
      - image: redis
        name: redis
`)
	deploy := &v1beta1.Deployment{}

	j, err := yaml.YAMLToJSON(byt)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(j, &deploy) // here give your json file's byte array and address of your struct
	if err != nil {
		panic(err)
	}

	res := schema.Validate(deploy)
	if res != nil {
		fmt.Println(res)
	}
}
