package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Animal struct {
	Animal  []string
	Name    string
	Secrets []map[string]string
}

func main() {
	viper.SetConfigName("demo")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}

	var animal Animal
	err = viper.Unmarshal(&animal)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(animal.Name)
}
