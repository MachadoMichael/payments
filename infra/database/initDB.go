package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MachadoMichael/payments/infra"
	"github.com/go-redis/redis/v8"
)

var StoreRepo *Repo
var PaymentRepo *Repo
var storeClient *redis.Client
var paymentClient *redis.Client

func Init() {
	ctx := context.Background()

	rdbStore := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DATABASE_ADDRESS"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DB:       infra.Config.DbStores,
	})

	rdbPayment := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DATABASE_ADDRESS"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DB:       infra.Config.DbPayments,
	})

	pong, err := rdbStore.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	pong2, err := rdbStore.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(pong)
	fmt.Printf(pong2)

	storeClient = rdbStore
	paymentClient = rdbPayment
	StoreRepo = NewRepo(ctx, rdbStore)
	PaymentRepo = NewRepo(ctx, rdbPayment)
}

func CloseDbStore() {
	storeClient.Close()
}

func CloseDbPayment() {
	paymentClient.Close()
}
