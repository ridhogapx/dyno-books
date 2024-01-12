package helper

import (
	"fmt"
	"log"

	"dyno-books/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Declare table schema
type Data struct {
  UserEmail string `json:"user_email"`
  // Relation one-to-many. Because single user can have lot of books.
  Book []Book
}


type Book struct {
  Title string `json:"title"`
  Author string `json:"author"`
}

func AddNote(data *Data) {
  av, err := dynamodbattribute.MarshalMap(data)

  if err != nil {
    log.Fatal(err)
  }

  input := &dynamodb.PutItemInput{
    Item: av,
    TableName: aws.String("books"),
  }

  _, err = config.Dyno.PutItem(input)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Successfully add data")
  

}
