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
	once.Do(
		func() {
			ConfigSingleton = new(Config)
			ConfigSingleton.Init()

			DynamoDBSingleton = new(adapters.DynamoDB)
			DynamoDBSingleton.Init(ConfigSingleton.Stage)
		})
}
