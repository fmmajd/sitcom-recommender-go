package shows

import "encoding/json"

type Episode struct {
	Season uint
	EpisodeNo uint `json:"episode"`
	Title string `json:"name"`
}

func (e Episode) MarshalJSON() ([]byte, error) {
	encoded := map[string]interface{}{
		"Season": e.Season,
		"EpisodeNo": e.EpisodeNo,
		"Title": e.Title,
	}
	return json.Marshal(encoded)
}
