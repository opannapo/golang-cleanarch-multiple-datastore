package app

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

//SetupConfig --> Function Setup Config
func SetupConfig(arg string) {
	var configFile string
	switch arg {
	case "local":
		configFile = "./config/config.json"
	case "dev":
		configFile = "./config/config-dev.json"
	case "prod":
		configFile = "./config/config-prod.json"
	}

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	println("Config Env : " + arg)
	println("Config File : " + configFile)
	println("Config Port : " + viper.GetString("server.address"))
}

//SetupDbConnection --> Function to Setup DB Connection
func SetupDbConnection() *gorm.DB {
	println("setupDbConnection")
	user := viper.GetString("database.user")
	password := viper.GetString("database.pass")
	dbname := viper.GetString("database.name")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open("mysql", connectionString)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}

	return db
}

//SetupCacheRedisConnection setup redis config
func SetupCacheRedisConnection() *redis.Client {
	addr := viper.GetString("cache_redis.addr")
	password := viper.GetString("cache_redis.password")
	db := viper.GetInt("cache_redis.db")
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return redisClient
}

//SetupFirestore --> Fuction to setup Firestore Configuration
func SetupFirestore() *firestore.Client {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./config/firebase/firebase-config.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
