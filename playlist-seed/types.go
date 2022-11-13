package playlist_seed

type Playlist struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Videos []Videos `json:"videos"`
}

type Videos struct {
	Id string `json:"id"`
}
