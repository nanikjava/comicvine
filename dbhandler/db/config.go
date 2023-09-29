package db

// We want the DB handler code to be stateless. Meaning
// it need to run and get the data and will not contain
// any logic that requires it to store some states.
// .....
// For example - The DB handler should just store the
// data provided to it as parameter and store it in the
// location it knows

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

}
