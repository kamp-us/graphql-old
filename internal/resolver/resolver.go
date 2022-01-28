package resolver

import "github.com/kamp-us/graphql/internal/clients"

// The QueryResolver is the entry point for all top-level read operations.
type QueryResolver struct {
	Clients *clients.Clients
}

func NewRoot(c *clients.Clients) (*QueryResolver, error) {
	return &QueryResolver{
		Clients: c,
	}, nil
}
