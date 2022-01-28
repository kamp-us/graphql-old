package clients

import (
	"net/http"
	"os"

	courseapi "github.com/kamp-us/course-api/rpc/course-api"
)

type Clients struct {
	CourseAPI courseapi.CourseAPI
}

func NewClients() (*Clients, error) {
	return &Clients{
		CourseAPI: courseapi.NewCourseAPIProtobufClient(os.Getenv("COURSEAPI_HOST"), &http.Client{}),
	}, nil
}
