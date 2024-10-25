package database

import (
	"context"
	"log"

	"github.com/vitor-chaves-lima/stop/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBManager struct {
	Database *mongo.Database
}

func NewMongoDBManager(ctx context.Context, config config.MongoDBConfig) *MongoDBManager {
	mongoDsn := config.GenerateDSN()
	clientOptions := options.Client().ApplyURI(mongoDsn)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return &MongoDBManager{Database: client.Database(config.Database)}
}

func (m *MongoDBManager) Disconnect() {
	_ = m.Database.Client().Disconnect(context.Background())
}
