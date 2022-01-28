package loader

import (
	"context"

	"github.com/graph-gophers/dataloader"
	courseapi "github.com/kamp-us/course-api/rpc/course-api"
)

const (
	courseByID key = "courseByID"
	// courseBySlug key = "courseBySlug"
)

type CourseGetter interface {
	GetBatchCourses(context.Context, *courseapi.GetBatchCoursesRequest) (*courseapi.GetBatchCoursesResponse, error)
}

type CourseLoader struct {
	Client CourseGetter
}

func newCourseByIDLoader(client CourseGetter) dataloader.BatchFunc {
	return CourseLoader{Client: client}.loadBatch
}

func (cl CourseLoader) loadBatch(ctx context.Context, ids dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(ids)
		results = make([]*dataloader.Result, n)
		reqIds  = make([]string, n)
	)

	for i, id := range ids {
		reqIds[i] = id.String()
	}

	resp, err := cl.Client.GetBatchCourses(ctx, &courseapi.GetBatchCoursesRequest{
		Ids: reqIds,
	})

	if err != nil {
		for i := range ids {
			results[i] = &dataloader.Result{Error: err}
		}
	} else {
		for i := range ids {
			results[i] = &dataloader.Result{Data: resp.Courses[i]}
		}
	}

	return results
}
