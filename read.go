package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

const baseURL = "https://trends.google.com/trends/trendingsearches/daily/rss?geo=%s"

func readGoogleTrends(country string) (result []string, err error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(fmt.Sprintf(baseURL, country))
	if err != nil {
		return nil, err
	}
	fmt.Println(feed.Title)

	for i, item := range feed.Items {
		result = append(result, fmt.Sprintf("%d: %s", i+1, item.Title))
	}

	return result, nil
}
