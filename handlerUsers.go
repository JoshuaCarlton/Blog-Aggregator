package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	ctx := context.Background()

	names, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, name := range names {
		fmt.Printf("* %s", name)
		if name == s.config.CurrentUserName {
			fmt.Print(" (current)\n")
		} else {
			fmt.Print("\n")
		}
	}
	return nil
}
