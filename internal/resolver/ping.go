package resolver

type PingResponse struct {
	message string
}

func (r *PingResponse) Message() *string {
	return &r.message
}

func (q *QueryResolver) Ping() (*PingResponse, error) {
	pong := "pong"
	return &PingResponse{
		message: pong,
	}, nil
}
