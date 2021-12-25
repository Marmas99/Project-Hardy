package database

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
)

type ItemToDelete struct {
	Title string `json:"title"`
}

const tablename = "Todo-Table"

func DeleteItem(itemData []byte, apilogger *logrus.Logger) {
	var itemToDelete ItemToDelete

	err := json.Unmarshal(itemData, &itemToDelete)
	if err != nil {
		apilogger.Errorln(fmt.Sprintln("Error while unmarshaling itemData: %s", err.Error()))
	}
	dynamodbClient := GetDynamoDbClient()

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"title": {
				S: aws.String(itemToDelete.Title),
			},
		},
		TableName: aws.String(tableName),
	}

	apilogger.Infoln(fmt.Sprintf("Attempting to delete item: %s", itemToDelete.Title))
	_, err = dynamodbClient.DeleteItem(input)
	if err != nil {
		apilogger.Errorln(fmt.Sprintf("Error while deleting item: %s, message: %s", itemToDelete.Title, err.Error()))
	} else {
		apilogger.Infoln(fmt.Sprintf("Successfully deleted item: %s", itemToDelete.Title))
	}
}
