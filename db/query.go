package db

import (
	"context"
	"fmt"
)

func CreateTest() {
	result, err := Mongo.DB.Collection("soominTest").
		InsertOne(context.TODO(), map[string]any{
			"key1": "value1",
			"key2": map[string]any{"subKey1": "subValue1"},
		})

	fmt.Println(result, err)
}
