package bootstrap

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Bootstrap struct {
	MgClient *mongo.Client
}

// New :
func New(ctx context.Context) *Bootstrap {
	bs := new(Bootstrap)

	bs.InitMongoDB(ctx)

	return bs
}
