package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"tnals5152.com/api-gateway/model"
)

func CreateTest() {
	result, err := Mongo.DB.Collection("soominTest").
		InsertOne(context.TODO(), map[string]any{
			"key1": "value1",
			"key2": map[string]any{"subKey1": "subValue1"},
		})

	fmt.Println(result, err)
}

func GetTest() {
	var data []model.PathInfo
	res, err := Mongo.DB.Collection("soominTest").Find(context.TODO(), bson.M{})

	res.All(context.TODO(), &data)

	fmt.Println(res, err)
}
