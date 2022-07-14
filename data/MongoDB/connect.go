package DatabaseMongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"	
)



func Init(MongoDBURL string, MongoDBMinPoolSize uint64)  (*mongo.Client, error) {
	MongoClient, err := mongo.Connect(
		context.Background(), 
		options.Client().ApplyURI(MongoDBURL).SetMinPoolSize(MongoDBMinPoolSize),
	)
	if err != nil {
		return nil, err
	}

	if err := MongoClient.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	return  MongoClient, nil
}

func Disconnect(MongoClient *mongo.Client) {
	if err := MongoClient.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}