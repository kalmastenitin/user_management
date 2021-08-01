package config

import (
	"context"
	"log"
	"time"

	"github.com/kalmastenitin/user_management/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	// database configuration
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// close the connection at the end
	defer cancel()

	if err != nil {
		log.Fatal("cannot connect to database", err)
	} else {
		log.Println("connected to database!")
	}

	db := client.Database("usergo")

	models.GroupCollection(db)
	models.GroupPermissionCollection(db)
	models.PermissionCollection(db)
	models.RoleCollection(db)
	models.UserCollection(db)
	return

}
