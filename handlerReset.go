package main

import (
	"context"
)

func handlerReset(s *state, _ command) error {
	ctx := context.Background()

	err := s.db.ResetTable(ctx)
	return err
}
