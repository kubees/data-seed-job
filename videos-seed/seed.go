package videos_seed

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"os"
)

func SeedVideosData(client *redis.Client, ctx context.Context) {
	videos, err := getVideosFromJson()
	if err != nil {
		return
	}
	for _, video := range videos {
		videoJson, err := json.Marshal(video)
		if err != nil {
			fmt.Println("Error while marshaling")
			return
		}
		err = client.Set(ctx, video.Id, videoJson, 0).Err()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func getVideosFromJson() ([]Videos, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("videos-seed/videos.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return []Videos{}, err
	}
	fmt.Println("Successfully Opened videos.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var videos []Videos

	err = json.Unmarshal(byteValue, &videos)
	if err != nil {
		fmt.Println(err)
		return []Videos{}, err
	}
	return videos, nil
}
