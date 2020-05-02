package shows

import (
	"encoding/json"
	"io/ioutil"
)

func LoadShowByTitle(title string) (*TvShow, error) {
	show := TvShow{}
	file, err := ioutil.ReadFile(dataFilePathByTitle(title))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &show)
	if err != nil {
		return nil, err
	}

	return &show, nil
}
