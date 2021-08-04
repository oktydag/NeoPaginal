package mongoClient

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
	"neopaginal-passanger/errors"
	"time"
)

func newDatabase(uri, databaseName string, dbOpts ...*options.DatabaseOptions) (db *mongo.Database, err error) {

	clientOptions := options.
		Client().
		ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return db, errors.NewWithCause(ConnectionError, err, uri)
	}

	ctxWithTimeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctxWithTimeout)
	if err != nil {
		return db, errors.NewWithCause(ConnectionError, err, uri)
	}

	db = client.Database(databaseName, dbOpts...)

	return db, err
}

func NewDatabaseWithSecondary(uri, databaseName string) (db *mongo.Database, err error) {
	secondary := readpref.SecondaryPreferred()
	dbOpts := options.Database().SetReadPreference(secondary)
	return newDatabase(uri, databaseName, dbOpts)
}
