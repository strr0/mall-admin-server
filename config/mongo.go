package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	ctx     context.Context
	cancel  context.CancelFunc
	mongodb *mongo.Client
)

type MongoConfig struct {
	Url      string
	Username string
	Password string
}

func InitMongo() {
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var mongoConfig MongoConfig
	err := settings.UnmarshalKey("mongo", &mongoConfig)
	if err != nil {
		log.Fatalf("get config failed: %v", err)
	}
	uri := fmt.Sprintf("mongodb://%s:%s@%s", mongoConfig.Username, mongoConfig.Password, mongoConfig.Url)
	mongodb, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("connect mongo failed: %v", err)
	}
}

func CLoseMongo() {
	_ = mongodb.Disconnect(ctx)
	cancel()
}

func GetMongo() *mongo.Client {
	return mongodb
}
