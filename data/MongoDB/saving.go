package DatabaseMongoDB

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"	
	"go.mongodb.org/mongo-driver/mongo/options"	
)

func BunchSaving(msgC chan string, collection *mongo.Collection) {
	MSet := []interface{}{}
	opts := options.InsertMany().SetOrdered(false)
	for {
		select {
		case thisMsg :=<- msgC:
			// Decode thisMsg
			ParsedMsg := []byte(thisMsg)
			var thisData bson.M
			if err := bson.UnmarshalExtJSON(ParsedMsg, true, &thisData); err != nil {
				fmt.Println(err)
				continue
			}

			MSet = append(MSet, thisData)

			if len(MSet) >= 10 {
				if _, err := collection.InsertMany(context.Background(), MSet, opts); err != nil {
					fmt.Println(err)
				} else {
					MSet = []interface{}{}
				}
			} else {
				continue
			}
		}
	}
}
