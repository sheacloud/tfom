package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/tfom/internal/graph/generated"
	"github.com/sheacloud/tfom/pkg/models"
)

// Versions is the resolver for the versions field.
func (r *moduleGroupResolver) Versions(ctx context.Context, obj *models.ModuleGroup, limit *int, nextCursor *string) (*models.ModuleVersions, error) {
	if limit == nil {
		limit = aws.Int(100)
	}
	return r.apiClient.GetModuleVersions(ctx, obj.ModuleGroupId, int32(*limit), aws.ToString(nextCursor))
}

// CreateModuleGroup is the resolver for the createModuleGroup field.
func (r *mutationResolver) CreateModuleGroup(ctx context.Context, moduleGroup models.NewModuleGroup) (*models.ModuleGroup, error) {
	return r.apiClient.PutModuleGroup(ctx, &moduleGroup)
}

// DeleteModuleGroup is the resolver for the deleteModuleGroup field.
func (r *mutationResolver) DeleteModuleGroup(ctx context.Context, moduleGroupID string) (bool, error) {
	err := r.apiClient.DeleteModuleGroup(ctx, moduleGroupID)
	return err == nil, err
}

// ModuleGroup is the resolver for the moduleGroup field.
func (r *queryResolver) ModuleGroup(ctx context.Context, moduleGroupID string) (*models.ModuleGroup, error) {
	return r.apiClient.GetModuleGroup(ctx, moduleGroupID)
}

// ModuleGroups is the resolver for the moduleGroups field.
func (r *queryResolver) ModuleGroups(ctx context.Context, limit *int, nextCursor *string) (*models.ModuleGroups, error) {
	if limit == nil {
		limit = aws.Int(100)
	}
	return r.apiClient.GetModuleGroups(ctx, int32(*limit), aws.ToString(nextCursor))
}

// ModuleGroup returns generated.ModuleGroupResolver implementation.
func (r *Resolver) ModuleGroup() generated.ModuleGroupResolver { return &moduleGroupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type moduleGroupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
