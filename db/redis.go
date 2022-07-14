package db

import (
	"context"
	"math/rand"
	"url_shortener_main/encoding"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	Client *redis.Client
}

func NewDB(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}

func (db *Database) EncodeAndSaveLink(link string) (string, error) {
	encodedLink := encoding.Encode(rand.Uint64()) // generate random string (Uint64 --> Base58 str)
	for used := db.CheckExists(encodedLink); used == true; used = db.CheckExists(encodedLink) {
		encodedLink = encoding.Encode(rand.Uint64()) // regenerate if we got collision
	}
	err := db.Client.Set(context.Background(), encodedLink, link, 0) // no expiration date for now
	return encodedLink, err.Err()
}

func (db *Database) CheckExists(encodedLink string) bool {
	_, err := db.Client.Get(context.Background(), encodedLink).Result()
	if err == redis.Nil {
		return false
	}
	if err != nil {
		panic(err)
	}
	return true
}

func (db *Database) DecodeLink(encodedLink string) (string, error) {
	res, err := db.Client.Get(context.Background(), encodedLink).Result()
	return res, err
}
