package singleton

import (
	"sync"
	"ukiyo/adapters"
	"ukiyo/configuration"
)

type Singleton struct{}

var once sync.Once

func (s *Singleton) Init() {
	once.Do(
		func() {
			configuration.ConfigurationSingleton.Init()
			adapters.DynamoDBSingleton.Init()
		})
}
