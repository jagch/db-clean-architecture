package main

import (
	"fmt"
	"log"

	"github.com/jagch/db-clean-architecture/db"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	/*uri := "mongodb+srv://userMongodb:passMongodb@cluster0.bobtd.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"*/
	protocol := viper.GetString("dbs.mongodb.protocol")
	user := viper.GetString("dbs.mongodb.credentials.user")
	pass := viper.GetString("dbs.mongodb.credentials.pass")
	host := viper.GetString("dbs.mongodb.host")
	retryWrites := viper.GetString("dbs.mongodb.options.retryWrites")
	w := viper.GetString("dbs.mongodb.options.w")
	uri := fmt.Sprintf("%s://%s:%s@%s/myFirstDatabase?retryWrites=%s&w=%s", protocol, user, pass, host, retryWrites, w)
	timeout := viper.GetInt("context.timeout")
	msgConnected := viper.GetString("dbs.mongodb.msg.connected")

	db := new(db.Mongodb)
	db.Init(uri, timeout, msgConnected)
	client, ctx, cancel, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = db.Close(client, ctx, cancel)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Ping mongoDB with Ping method
	err = db.Ping(client, ctx)
	if err != nil {
		log.Fatal(err)
	}

}
