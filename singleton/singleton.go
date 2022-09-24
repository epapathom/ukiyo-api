package singleton

import (
	"sync"
	"ukiyo/adapters"
	"ukiyo/business"
	"ukiyo/configuration"
)

type Singleton struct{}

var once sync.Once

func (s *Singleton) Init() {
	once.Do(
		func() {
			configuration.ConfigurationSingleton = new(configuration.Configuration)
			configuration.ConfigurationSingleton.Init()

			adapters.DynamoDBSingleton = new(adapters.DynamoDB)
			adapters.DynamoDBSingleton.Init()

			business.LogicSingleton = new(business.Logic)
		})
}
