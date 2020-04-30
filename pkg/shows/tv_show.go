package shows

type TvShow struct {
	Name string
	episodateId uint
	Description string
	Status string
	Runtime uint
	Genres []string
	Episodes []Episode
}
