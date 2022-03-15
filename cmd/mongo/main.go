package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&ssl=false"))
	if err != nil {
		panic(err)
	}

	col := mc.Database("coolcar").Collection("account")
	res := col.FindOne(c, bson.M{
		"open_id":	"V",
	})

	var row struct{
		ID 		primitive.ObjectID 	`bson:"_id"`
		OpenID 	string				`bson:"open_id"`
		Money 	int					`bson:"money"`
	}

	err = res.Decode(&row)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", row)
}
