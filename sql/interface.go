package sql

import "net/url"

type Storage interface {
	// Insert logs a new reading record.
	Insert(*url.URL) error
	// Select gets the url with the lowest retrievability of memory.
	Select() (*url.URL, error)
	// Delete will remove an url permanently from the storage.
	Delete(*url.URL) error
}
