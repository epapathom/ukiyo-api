package adapters

import (
	"context"
	"log"
	"ukiyo/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDB struct {
	Client *dynamodb.Client
}

func (ddb *DynamoDB) Init() {
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	ddb.Client = dynamodb.NewFromConfig(config)
}

func (ddb *DynamoDB) GetPaintingByName(japaneseName string, englishName string) models.Painting {
	var key map[string]types.AttributeValue
	var painting models.Painting

	if japaneseName != "" {
		key = map[string]types.AttributeValue{
			"NameJapanese": &types.AttributeValueMemberS{Value: japaneseName},
		}
	} else if englishName != "" {
		key = map[string]types.AttributeValue{
			"NameEnglish": &types.AttributeValueMemberS{Value: englishName},
		}
	}

	getItemInput := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String("Paintings"),
	}

	output, err := ddb.Client.GetItem(context.TODO(), getItemInput)
	if err != nil {
		log.Fatalf("unable to get video game, %v", err)
	}

	err = attributevalue.UnmarshalMap(output.Item, &painting)
	if err != nil {
		log.Fatalf("unable to unmarshal map, %v", err)
	}

	return painting
}
