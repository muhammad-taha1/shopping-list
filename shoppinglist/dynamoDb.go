package shoppinglist

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

var svc *dynamodb.DynamoDB
var tableName = "ItemCollection"

func InitSession() {
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	// 	SharedConfigState: session.SharedConfigEnable,
	// }))

	
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("ca-central-1"),
		Endpoint: aws.String("http://localhost:8000")}))

	// Create DynamoDB client
	svc = dynamodb.New(sess)

	_, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// log.Println("Tables:")
	// log.Println("")

	// for _, n := range result.TableNames {
	// 	log.Println(*n)
	// }
}

func addNewItem(item CartItem) {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Println("Got error marshalling new item:")
		log.Println(err.Error())
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Println("Got error when Putting Item:")
		log.Println(err.Error())
		os.Exit(1)
	}

	log.Println("Successfully added new item")
}

func getItemsForUser(user string) (items []CartItem) {

	filt := expression.Name("User").Equal(expression.Value(user))

	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(tableName),
	}

	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	if len(items) < 1 {
		fmt.Println("Could not find any item for user: " + user)
		return
	}

	fmt.Println("Found items for user: " + user)
	return
}
