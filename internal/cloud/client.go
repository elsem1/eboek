package cloud

import "eboek/internal/storage"

func NewClient(store *storage.Store) *Client { // ontvangt een een store
	return &Client{}
}

type Client struct{}
