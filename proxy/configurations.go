package proxy

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadServiceConfigs(service_name string) *Service{
	var services []Service
	fileData, err := ioutil.ReadFile("./configs/proxyconfs.json")
	if err != nil{
		log.Fatal(err)
	}
	err = json.Unmarshal(fileData, &services)
	if err != nil{
		log.Fatal(err)
	}

	for i := range services{
		if services[i].ServiceName == service_name{
			return &services[i] 
		}
	}

	return nil
}