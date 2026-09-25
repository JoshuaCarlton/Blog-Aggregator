package main

import (
	"github.com/JoshuaCarlton/gator/internal/config"
	"github.com/JoshuaCarlton/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
