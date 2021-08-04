package mongoClient

import (
	goErrors "errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"golang.org/x/net/context"
	"neopaginal-command/errors"
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

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return db, errors.NewWithCause(PingError, err, uri)
	}

	db = client.Database(databaseName, dbOpts...)

	return db, err
}

func NewDatabase(uri, databaseName string) (db *mongo.Database, err error) {
	return newDatabase(uri, databaseName)
}

func NewDatabaseWithSecondary(uri, databaseName string) (db *mongo.Database, err error) {
	secondary := readpref.SecondaryPreferred()
	dbOpts := options.Database().SetReadPreference(secondary)
	return newDatabase(uri, databaseName, dbOpts)
}

func IsDuplicateError(err error) bool {
	var e mongo.WriteException
	if goErrors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}
