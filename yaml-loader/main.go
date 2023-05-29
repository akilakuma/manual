package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"fmt"
)

type DBConf struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Conf struct {
	DB map[string]DBConf `yaml:"mysql"`
}

func main() {


	var c Conf


	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(c)
}
