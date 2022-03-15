package dao

import (
	"context"
	mgo "coolcar/shared/mongo"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDField = "open_id"

// Mongo defines a dao(Data Access Object).
type Mongo struct {
	col *mongo.Collection
}

// NewMongo 是构造方法，返回一个Mongo结构体，通向db中名为account的collection
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

// ResolveAccountID 由用户唯一标识openID得到accountID
func (m *Mongo) ResolveAccountID(c context.Context, openID string) (string, error) {
	res := m.col.FindOneAndUpdate(
		c,
		bson.M{
			openIDField: openID,
		},
		mgo.Set(bson.M{
			openIDField: openID,
		}),
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)
	if res.Err() != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate: %v", res.Err())
	}

	var row mgo.ObjID

	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("cannot decode result: %v", res.Err())
	}

	return row.ID.Hex(), nil
}

