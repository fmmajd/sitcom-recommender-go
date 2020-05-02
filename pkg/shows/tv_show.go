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
	return dataFilePathByTitle(s.Name)
}

func dataFilePathByTitle(title string) string {
	wd, _ := os.Getwd()
	return wd+"/data/shows/"+title+".json"
}