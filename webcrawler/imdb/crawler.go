package imdb

import (
	"fmt"
	"sync"

	"github.com/gocolly/colly"
)

const moviesPerPage = 50

// Crawler is responsible for retrieving information from the top 1000 movies from IMDB
type Crawler struct{}

// Start the crawler
func (crawler *Crawler) Start(movieChannel chan Movie, doneChannel chan bool) {
	var wg sync.WaitGroup
	for moviePosition := 1; moviePosition < 1000; moviePosition += moviesPerPage {
		wg.Add(1)
		go crawler.crawlMovieRankingPage(moviePosition, movieChannel, &wg)
	}

	wg.Wait()
	doneChannel <- true
}

// crawlMovieRankingPage extracts all movies from the top 1000 movies ranking.
// The startingPosition parameter represents the first movie that will be displayed on that page. By
// using it, we can iterate over all movies in the ranking just like a user would do by using the navigation
// buttons.
func (crawler *Crawler) crawlMovieRankingPage(startingPosition int, movieChannel chan Movie, wg *sync.WaitGroup) {
	collector := colly.NewCollector()

	collector.OnHTML(".lister-list", func(listElement *colly.HTMLElement) {
		listElement.ForEach(".lister-item.mode-simple", func(index int, movieElement *colly.HTMLElement) {
			movie := Movie{}
			movie.URL = movieElement.ChildAttr(".lister-item-header a", "href")
			movie.Name = movieElement.ChildText(".lister-item-header a")
			movie.Cast = movieElement.ChildAttr(".lister-item-header span:nth-child(2)", "title")

			movieChannel <- movie
		})

		wg.Done()
	})

	pageURL := fmt.Sprintf("https://www.imdb.com/search/title/?groups=top_1000&view=simple&sort=user_rating,desc&start=%d&ref_=adv_nxt", startingPosition)
	collector.Visit(pageURL)
}

// NewCrawler creates a new instance of a IMDB crawler
func NewCrawler() *Crawler {
	return &Crawler{}
}
