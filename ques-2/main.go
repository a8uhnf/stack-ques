package main

import (
    "log"

    "gopkg.in/yaml.v2"
)

type Data struct {
    Entry []Entry `yaml:"entries"`
}

type Entry map[string]string

var dat string = `
entries: 
  - keya1: val1
    keya2: val2
  - keyb1: val1
    keyb2: val2
  - val3: ""`

func main() {
    out := Data{}
    if err := yaml.Unmarshal([]byte(dat), &out); err != nil {
        log.Fatal(err)
    }

    log.Printf("%+v", out)
}