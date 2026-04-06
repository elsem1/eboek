package sync

import (
	"eboek/internal/cloud"
	"eboek/internal/storage"
)

func NewEngine(client *cloud.Client, store *storage.Store) *Engine { // ontvangt een client en een store
	return &Engine{}
}

type Engine struct{}
