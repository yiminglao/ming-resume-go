package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go_resume/ent"
	"go_resume/graph/generated"
	"go_resume/graph/model"
)

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UserInput) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddEdu(ctx context.Context, input *model.EduInput) (*model.Edu, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateEdu(ctx context.Context, input *model.EduInput) (*model.Edu, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelEdu(ctx context.Context, id int) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetEduByID(ctx context.Context, id int) (*model.Edu, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllEdu(ctx context.Context) ([]*model.Edu, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetExperienceByID(ctx context.Context, id int) (*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllExp(ctx context.Context) ([]*model.Experience, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetMessageByID(ctx context.Context, id int) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllMessage(ctx context.Context) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetProjectByID(ctx context.Context, id int) (*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllProject(ctx context.Context) ([]*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
