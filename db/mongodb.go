package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongodb struct {
	uri          string
	timeout      int
	msgConnected string
}

func (mdb *Mongodb) Init(uri string, timeout int, msgConnected string) {
	mdb.uri = uri
	mdb.timeout = timeout
	mdb.msgConnected = msgConnected
}

func (mdb *Mongodb) Connect() (client interface{}, ctx context.Context, cancel context.CancelFunc, err error) {
	ctx, cancel = context.WithTimeout(context.Background(), time.Duration(mdb.timeout)*time.Second)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mdb.uri))
	return
}

func (mdb *Mongodb) Close(clientI interface{}, ctx context.Context, cancel context.CancelFunc) (err error) {

	//parse to *mongo.Client
	client := mdb.parseToClient(clientI)

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err = client.Disconnect(ctx); err != nil {
			return
		}
	}()
	return
}

func (mdb *Mongodb) Ping(clientI interface{}, ctx context.Context) (err error) {

	//parse to *mongo.Client
	client := mdb.parseToClient(clientI)

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return
	}
	fmt.Println(mdb.msgConnected)
	return
}

func (mdb *Mongodb) parseToClient(client interface{}) *mongo.Client {
	return client.(*mongo.Client)
}

//uri := "mongodb+srv://userMongodb:passMongodb>@cluster0.bobtd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
