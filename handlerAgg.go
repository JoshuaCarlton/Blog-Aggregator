package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.arguments) < 1 {
		return fmt.Errorf("you must provide the time to wait between requests")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}
