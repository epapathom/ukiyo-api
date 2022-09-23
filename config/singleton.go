package config

import (
	"sync"
	"ukiyo/adapters"
)

var once sync.Once

type Singleton struct{}

var ConfigSingleton *Config
var DynamoDBSingleton *adapters.DynamoDB

func (s *Singleton) Init() {
	if ConfigSingleton == nil {
		once.Do(
			func() {
				ConfigSingleton = new(Config)
				ConfigSingleton.Init()
			})
	}

	if DynamoDBSingleton == nil {
		once.Do(
			func() {
				DynamoDBSingleton = new(adapters.DynamoDB)
				DynamoDBSingleton.Init()
			})
	}
}
