package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	playlist "github.com/kubees/data-seed-job/playlist-seed"
	videos "github.com/kubees/data-seed-job/videos-seed"
	"os"
)

var microservice = os.Getenv("MICROSERVICE")
var redisHost = os.Getenv("REDIS_HOST")
var redisPort = os.Getenv("REDIS_PORT")
var password = os.Getenv("PASSWORD")
var client *redis.Client

func main() {
	r := redis.NewClient(&redis.Options{
		Addr: redisHost + ":" + redisPort,
		DB:   0,
		Password: password,
	})
	client = r
	ctx := context.Background()

	if microservice == "playlist" {
		playlist.SeedPlaylistsData(client, ctx)
	} else if microservice == "videos" {
		videos.SeedVideosData(client, ctx)
	}
}
