package database

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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

func getDynamoDbClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	return svc

}

func InsertNewItem(itemData []byte, writer http.ResponseWriter, apilogger *logrus.Logger) {
	var item Item

	json.Unmarshal(itemData, &item)
	dynamoClient := getDynamoDbClient()

	apilogger.Infoln("Marshaling new Item to be inserted")
	marshaldItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		apilogger.Error("Error while marshaling new Item")
	}

	input := &dynamodb.PutItemInput{
		Item:      marshaldItem,
		TableName: aws.String(tableName),
	}

	apilogger.Info(fmt.Sprintf("Putting item in table: %s", tableName))
	_, err = dynamoClient.PutItem(input)
	if err != nil {
		apilogger.Error(fmt.Sprintf("Error while putting item in table: %s", err.Error()))
	} else {
		apilogger.Infoln(fmt.Sprintf("Item successfully put in table: %s", tableName))
	}
}
