package playlist_seed

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io/ioutil"
	"os"
)

func SeedPlaylistsData(client *redis.Client, ctx context.Context) {
	playlists, err := getPlaylistsFromJson()
	if err != nil {
		return
	}
	err = client.Set(ctx, "playlists", playlists, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getPlaylistsFromJson() ([]Playlist, error) {
	// Open our jsonFile
	jsonFile, err := os.Open("playlist-seed/playlists.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return []Playlist{}, err
	}
	fmt.Println("Successfully Opened playlists.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var playlists []Playlist

	err = json.Unmarshal(byteValue, &playlists)
	if err != nil {
		fmt.Println(err)
		return []Playlist{}, err
	}
	return playlists, nil
}
