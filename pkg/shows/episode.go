package shows

import (
	"encoding/json"
)

type Episode struct {
	Season uint
	EpisodeNo uint
	Title string
}

func (e Episode) MarshalJSON() ([]byte, error) {
	encoded := map[string]interface{}{
		"Season": e.Season,
		"EpisodeNo": e.EpisodeNo,
		"Title": e.Title,
	}
	return json.Marshal(encoded)
}

func (e *Episode) UnmarshalJSON(b []byte) error {
	data := map[string]interface{}{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	var episodeNo, season float64
	var title string
	var exists bool

	episodeNo, exists = data["episode"].(float64)
	if !exists {
		episodeNo, exists = data["EpisodeNo"].(float64)
	}
	if exists {
		e.EpisodeNo = uint(episodeNo)
	}

	season, exists = data["season"].(float64)
	if !exists {
		season, exists = data["Season"].(float64)
	}
	if exists {
		e.Season = uint(season)
	}

	title, exists = data["name"].(string)
	if !exists {
		title, exists = data["Title"].(string)
	}
	if exists {
		e.Title = title
	}

	return nil
}
