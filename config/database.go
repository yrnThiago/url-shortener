package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Conn *mongo.Collection

func getDatabaseUrl() string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s.%s.mongodb.net/",
		Env.DBUsername,
		Env.DBPassword,
		Env.DBName,
		Env.DBIdk,
	)
}

func DatabaseInit() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(getDatabaseUrl()).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	//
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		Logger.Fatal("MongoDB did not pong!")
	}

	Logger.Info(
		"Successfully connected to MongoDB!",
	)

	Conn = client.Database(Env.DBName).Collection("urls")
}
