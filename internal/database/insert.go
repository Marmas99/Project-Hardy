package database

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sirupsen/logrus"
)

type Item struct {
	Title    string `json:"title"`
	Desc     string `json:"description"`
	Priority int    `json:"priority"`
}

const tableName = "Todo-Table"

func InsertNewItem(itemData []byte, apilogger *logrus.Logger) {
	var item Item

	json.Unmarshal(itemData, &item)
	dynamoClient := GetDynamoDbClient()

	marshaldItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		apilogger.Error("Error while marshaling new Item")
	}

	input := &dynamodb.PutItemInput{
		Item:      marshaldItem,
		TableName: aws.String(tableName),
	}

	apilogger.Infoln(fmt.Sprintf("Putting item in table: %s", tableName))
	_, err = dynamoClient.PutItem(input)
	if err != nil {
		apilogger.Error(fmt.Sprintf("Error while putting item in table: %s", err.Error()))
	} else {
		apilogger.Infoln(fmt.Sprintf("Item successfully put in table: %s", tableName))
	}
}
