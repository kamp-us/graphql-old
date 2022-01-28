package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	courseapi "github.com/kamp-us/course-api/rpc/course-api"
)

type CourseInput struct {
	ID   *graphql.ID
	Slug *string
}

type CourseArgs struct {
	// Input CourseInput
	ID   *graphql.ID
	Slug *string
}

type CourseResolver struct {
	Course *courseapi.Course
}

func (r *QueryResolver) Course(ctx context.Context, args *CourseArgs) (*CourseResolver, error) {
	course, err := r.Clients.CourseAPI.GetCourse(ctx, &courseapi.GetCourseRequest{
		ID: string(*args.ID),
	})

	if err != nil {
		return nil, err
	}

	return &CourseResolver{Course: course}, nil
}

func (r *CourseResolver) ID() graphql.ID {
	return graphql.ID(r.Course.ID)
}

func (r *CourseResolver) Slug() string {
	return r.Course.Slug
}

func (r *CourseResolver) UserID() graphql.ID {
	return graphql.ID(r.Course.UserId)
}

func (r *CourseResolver) Name() string {
	return r.Course.Name
}

func (r *CourseResolver) Description() string {
	return r.Course.Description
}
