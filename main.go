package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	playlist "github.com/kubees/data-seed-job/playlist-seed"
	"os"
)

var microservice = os.Getenv("MICROSERVICE")
var redisHost = os.Getenv("REDIS_HOST")
var redisPort = os.Getenv("REDIS_PORT")
var client *redis.Client

func main() {
	r := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
		DB:   0,
	})
	client = r
	ctx := context.Background()
	//if microservice == "playlist" {
	playlist.SeedPlaylistsData(client, ctx)
	//}
}
