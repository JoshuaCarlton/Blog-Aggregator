package main

import (
	"github.com/JoshuaCarlton/Blog-Aggregator/internal/config"
	"github.com/JoshuaCarlton/Blog-Aggregator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
