package db

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	ct "tnals5152.com/api-gateway/const"
)

func ConnectDB() {
	// ctx, cancel := context.WithTimeout(
	// 	context.Background(),
	// 	viper.GetDuration(ct.DBTimeout) * time.Second)
	// defer cancel()

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

	client, err := mongo.Connect(context.TODO(), clientOptions)

	fmt.Println(client, err)
}
