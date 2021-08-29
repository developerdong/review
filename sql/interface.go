package sql

import "net/url"

type Storage interface {
	// Add logs a new reading record.
	Add(*url.URL) error
	// Review gets the url with the lowest retrievability of memory.
	Review() (*url.URL, error)
	// Delete will remove an url permanently from the storage.
	Delete(*url.URL) error
}
