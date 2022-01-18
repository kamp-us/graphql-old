package resolver

// The QueryResolver is the entry point for all top-level read operations.
type QueryResolver struct {
}

func NewRoot() (*QueryResolver, error) {
	return &QueryResolver{}, nil
}
