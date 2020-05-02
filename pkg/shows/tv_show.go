package shows

import "os"

type TvShow struct {
	Name string
	episodateId uint
	Description string
	Status string
	Runtime uint
	Genres []string
	Episodes []Episode
}

func (s TvShow) dataFile() string {
	wd, _ := os.Getwd()
	return wd+"/data/shows/"+s.Name+".json"
}