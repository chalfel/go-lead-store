package db

import (
	"context"

	"github.com/chalfel/go-lead-store/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConnection struct {
	Conn *mongo.Client
}

func NewDatabaseConnection(c *config.DatabaseConfig) (*DatabaseConnection, error) {
	var err error
	conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI(c.Dns))

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{Conn: conn}, nil
}

func (d *DatabaseConnection) CloseConnection() {
	d.Conn.Disconnect(context.Background())
}
