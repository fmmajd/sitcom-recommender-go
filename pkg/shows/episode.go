package shows

type Episode struct {
	Season uint
	EpisodeNo uint `json:"episode"`
	Title string `json:"name"`
}
