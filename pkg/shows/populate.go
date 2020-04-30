package shows

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func PopulateTvShows() {
	for showTitle, episodateId := range includedShowsEpisodateIds {
		log.Printf("Starting to populate %s\n", showTitle)
		log.Println(episodateId)
		err, show := fetchEpisodateDataByShowId(episodateId)
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Printf("%+v\n", show)
		//log.Println(show)
		file, _ := json.MarshalIndent(show, "", " ")

		_ = ioutil.WriteFile("test.json", file, 0644)

		break
	}
}
