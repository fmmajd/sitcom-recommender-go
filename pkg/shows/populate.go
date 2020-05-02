package shows

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func PopulateTvShows() {
	for showTitle, episodateId := range includedShowsEpisodateIds {
		log.Printf("Populating %s\n...", showTitle)
		err, show := fetchEpisodateDataByShowId(episodateId)
		if err != nil {
			log.Fatalln(err)
		}

		err = writeShowToFile(*show)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func writeShowToFile(show TvShow) error{
	file, err := json.MarshalIndent(show, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(show.dataFile(), file, 0644)
	if err != nil {
		return err
	}

	return nil
}
