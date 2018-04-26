package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	config, err := yaml.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("source_path"))
	fmt.Println(config.Get("target_path"))
	newNode,err:=yaml.Child(config.Root,"servers")
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Printf("%+v",newNode)

}