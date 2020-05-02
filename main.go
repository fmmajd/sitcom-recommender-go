package main

import (
	"fmt"
	"github.com/fmmajd/sitcom-recommender-go/pkg/shows"
	"math/rand"
)

func main() {
	//shows.PopulateTvShows()

	fmt.Println("Select a Number: ")
	for i, title := range shows.AllShows {
		fmt.Printf("%d: %s\n", i+1, title)
	}
	var selectedIndex uint
	fmt.Scanln(&selectedIndex)
	title := shows.AllShows[selectedIndex-1]
	show,_ := shows.LoadShowByTitle(title)
	episodes := show.Episodes
	randInt := rand.Intn(len(episodes))
	episode := episodes[randInt-1]
	fmt.Printf("%+v\n", episode)
	fmt.Printf("Season %d, Episode %d: %s\n", episode.Season, episode.EpisodeNo, episode.Title)
}
