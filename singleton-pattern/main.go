package main

import "fmt"

type singletonCon struct {
	url  string
	port string
}

var singletonConInstances *singletonCon

func GetInstance() *singletonCon {
	if singletonConInstances == nil {

		fmt.Println("Creating singleton instance")

		singletonConInstances = &singletonCon{
			url:  "the url",
			port: "the port",
		}
	} else {
		fmt.Println("singleton already exist")
	}

	return singletonConInstances
}

func main() {
	for i := 0; i < 10; i++ {
		GetInstance()
	}
}
