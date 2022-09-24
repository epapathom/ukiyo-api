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

func (ddb *DynamoDB) GetPaintingsByArtist(artist string) []models.Painting {
	var paintings []models.Painting

	expression := map[string]types.AttributeValue{
		"Artist": &types.AttributeValueMemberS{Value: artist},
	}

	scanInput := &dynamodb.ScanInput{
		TableName:                 aws.String("Paintings"),
		ExpressionAttributeValues: expression,
	}

	output, err := ddb.Client.Scan(context.TODO(), scanInput)
	if err != nil {
		log.Fatalf("unable to get video game, %v", err)
	}

	err = attributevalue.UnmarshalListOfMaps(output.Items, &paintings)
	if err != nil {
		log.Fatalf("unable to unmarshal map, %v", err)
	}

	return paintings
}
