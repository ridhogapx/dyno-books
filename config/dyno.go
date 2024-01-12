package config

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var s = session.Must(session.NewSessionWithOptions(session.Options{
  SharedConfigState: session.SharedConfigEnable,
}))

var Dyno = dynamodb.New(s)
var TableName = "books"
