package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestMongo_ResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&ssl=false"))
	if err != nil {
		t.Errorf("cannot connect database: %v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	got, err := m.ResolveAccountID(c, "Jin")
	if err != nil {
		t.Errorf("cannot resolve account: %v", err)
	} else {
		wanted := "6218470b47b376ee0bdf05a4"
		if got != wanted {
			t.Errorf("got: %q, but want: %q", got, wanted)
		}
	}


}
