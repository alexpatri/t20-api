package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"t20-api/internal/utils/config"
)

var Database *mongo.Database

// Connect estabelece conexão com o Banco
func Connect(dbConfig config.DatabaseConfig) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(dbConfig.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Verificar conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(dbConfig.DatabaseName), nil
}

// Disconnect encerra a conexão com o Banco
func Disconnect(db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return db.Client().Disconnect(ctx)
}
