package shows

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func PopulateTvShows() {
	var tvShows []TvShow
	tmp := map[string]TvShow{}
	for showTitle, episodateId := range includedShowsEpisodateIds {
		log.Printf("Populating %s\n...", showTitle)
		err, show := fetchEpisodateDataByShowId(episodateId)
		if err != nil {
			log.Fatalln(err)
		}

		tvShows = append(tvShows, *show)
		tmp[showTitle] = *show
	}

	err := writeShowsToFile(tmp)
	if err != nil {
		log.Fatalln(err)
	}
}

func writeShowsToFile(shows map[string]TvShow) error{
	file, err := json.MarshalIndent(shows, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("shows_data.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}
