package bootstrap

import (
	"context"
	"log"
	"time"

	"github.com/moneylion-api/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (hdl *Bootstrap) InitMongoDB(ctx context.Context) {
	/*
	   Connect to my cluster
	*/
	client, err := mongo.NewClient(options.Client().ApplyURI(env.Config.Mongo_DB_Connection))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	hdl.MgClient = client

}
