package db

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db *dynamodb.DynamoDB

// Init dynamodb
func Init() {
	// os.Setenv("AWS_PROFILE", "ahnsv")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("/Users/rokaf/.aws/credentials.csv", ""),
	})
	if err != nil {
		log.Fatal(err)
	}
	db = dynamodb.New(sess)
}

func GetDB() *dynamodb.DynamoDB {
	return db
}
