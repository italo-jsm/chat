package config

import (
	"encoding/json"
	"io/ioutil"
	"sync"
)

var lock = &sync.Mutex{}

type Configuration struct {
    Database DatabaseConfiguration `json:"database"`
    Kafka KafkaConfiguration `json:"kafka"`
    NotifyEnabled bool
}

var configuration *Configuration

func GetInstance() *Configuration {
    if configuration == nil {
        lock.Lock()
        defer lock.Unlock()
        if configuration == nil {
            configuration = fillConfiguration()
        }
    }
    return configuration
}

func fillConfiguration() *Configuration{
    b, err := ioutil.ReadFile("conf.json")
    if(err != nil){
        panic(err.Error())
    }
    config := Configuration{}
    json.Unmarshal(b, &config)
	return &config
}