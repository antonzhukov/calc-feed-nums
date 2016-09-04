package service

import (
	"calc-feed-nums/app/domain"
	"calc-feed-nums/app/transport/dto"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const feedTimeout uint16 = 450 // in ms

// Calculates and returns numbers from feeds provided with get params
func CalculateNumbersFeed(w http.ResponseWriter, r *http.Request) {
	// Parse feed urls
	u, err := url.Parse(r.RequestURI)
	if err != nil {
		log.Fatal(err)
	}
	urls := u.Query()["u"]

	// Init
	var wg sync.WaitGroup
	wg.Add(len(urls))
	numbers := domain.NewNumbers()

	// Fill numbers struct with numbers from feed
	for _, feedUrl := range urls {
		go func(feedUrl string, numbers *domain.Numbers) {
			defer wg.Done()
			processFeed(feedUrl, numbers)
		}(feedUrl, &numbers)

	}
	wg.Wait()

	// Output numbers
	dtoNums := numbers.ConvertToTransfer()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dtoNums)

}

func processFeed(feedUrl string, numbers *domain.Numbers) {
	feed := dto.Numbers{}

	// Validate url format
	if !validateUrl(feedUrl) {
		err := errors.New(fmt.Sprintf("Invalid url: '%s'", feedUrl))
		log.Printf("%s", err)
		return
	}

	// Receive numbers feed
	err := getJson(feedUrl, &feed)
	if err != nil {
		err := errors.New(fmt.Sprintf("Failed to receive feed, url: '%s', error: '%s'", feedUrl, err))
		log.Printf("%s", err)
		return
	}

	// If there are any numbers in the feed, add to our numbers
	if len(feed.Numbers) > 0 {
		numbers.Add(feed)
	}
}

// Get content of url into predefined struct
func getJson(url string, target interface{}) error {
	// Init client
	timeout := time.Duration(time.Duration(feedTimeout) * time.Millisecond)
	c := http.Client{
		Timeout: timeout,
	}

	// Get contents
	r, err := c.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// Validate URL against obvious patterns
func validateUrl(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.ParseRequestURI(str)

	if err != nil {
		return false
	}
	if u.Host == "" || u.Scheme == "" {
		return false
	}

	return true
}
