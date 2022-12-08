package db

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	ct "tnals5152.com/api-gateway/const"
	"tnals5152.com/api-gateway/utils"
)

type MongoStruct struct {
	client *mongo.Client
	DB     *mongo.Database
}

var Mongo = &MongoStruct{}

func ConnectDB() {

	credentail := options.Credential{
		Username: viper.GetString(ct.DBUsername),
		Password: viper.GetString(ct.DBPassword),
	}

	clientOptions := options.Client().ApplyURI(
		"mongodb://" +
			viper.GetString(ct.DBHost) +
			":" +
			viper.GetString(ct.DBPort),
	).SetAuth(credentail)

	client, err := mongo.Connect(context.Background(), clientOptions)

	Mongo.client = client

	if err != nil {
		utils.Panic(err)
	}

	Mongo.DB = Mongo.client.Database(viper.GetString(ct.DBDatabase))
}
