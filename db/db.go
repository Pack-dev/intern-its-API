package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout = 30
)

// Resource :: referene to database
type Resource struct {
	DB *mongo.Database
}

func (r *Resource) Close() {
	ctx, cancel := InitContext()
	defer cancel()

	if err := r.DB.Client().Disconnect(ctx); err != nil {
		fmt.Println("Close connection falure, Something wrong...")
		return
	}

	fmt.Println("Close connection successfully")
}

func InitContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	return ctx, cancel
}

// CreateResource :: to create connection to database
func CreateResource() (*Resource, error) {
	_ = godotenv.Load()
	var admindb, dbName string

	if os.Getenv("MONGODB_ENDPOINT") != "" {
		admindb = os.Getenv("MONGODB_ENDPOINT")
	}
	if os.Getenv("MONGODB_NAME") != "" {
		dbName = os.Getenv("MONGODB_NAME")
	}

	fmt.Println("admindb : " + admindb)
	fmt.Println("dbName : " + dbName)
	client, err := mongo.NewClient(
		options.Client().ApplyURI(admindb),
		options.Client().SetMinPoolSize(1),
		options.Client().SetMaxPoolSize(1),
	)
	if err != nil {
		logrus.Errorf("Failed to create client: %v", err)
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		logrus.Errorf("Failed to connect to server: %v", err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		logrus.Errorf("Failed to ping cluster: %v", err)
		return nil, err
	}

	return &Resource{DB: client.Database(dbName)}, nil
}
