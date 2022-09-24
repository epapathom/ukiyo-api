package adapters

import (
	"context"
	"fmt"
	"log"
	"ukiyo/configuration"
	"ukiyo/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDB struct {
	Client    *dynamodb.Client
	TableName string
}

var DynamoDBSingleton *DynamoDB

func (ddb *DynamoDB) Init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	ddb.Client = dynamodb.NewFromConfig(cfg)
	ddb.TableName = fmt.Sprintf("Paintings-%s", configuration.ConfigurationSingleton.Stage)
}

func (ddb *DynamoDB) GetPaintingsByArtist(artist string) []models.Painting {
	var paintings []models.Painting

	tableName := aws.String(ddb.TableName)

	filterExpression := aws.String("contains(Artist, :Artist)")

	expressionAttributeValues := map[string]types.AttributeValue{
		":Artist": &types.AttributeValueMemberS{Value: artist},
	}

	scanInput := &dynamodb.ScanInput{
		TableName:                 tableName,
		FilterExpression:          filterExpression,
		ExpressionAttributeValues: expressionAttributeValues,
	}

	output, err := ddb.Client.Scan(context.TODO(), scanInput)
	if err != nil {
		log.Fatalf("unable to get paintings, %v", err)
	}

	err = attributevalue.UnmarshalListOfMaps(output.Items, &paintings)
	if err != nil {
		log.Fatalf("unable to unmarshal map, %v", err)
	}

	return paintings
}
