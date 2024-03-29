package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"

	"github.com/sheacloud/tfom/internal/graph/generated"
	"github.com/sheacloud/tfom/internal/terraform"
	"github.com/sheacloud/tfom/pkg/models"
)

// ModuleVersion is the resolver for the moduleVersion field.
func (r *moduleAssignmentResolver) ModuleVersion(ctx context.Context, obj *models.ModuleAssignment) (*models.ModuleVersion, error) {
	return r.apiClient.GetModuleVersionBatched(ctx, obj.ModuleVersionID)
}

// ModuleGroup is the resolver for the moduleGroup field.
func (r *moduleAssignmentResolver) ModuleGroup(ctx context.Context, obj *models.ModuleAssignment) (*models.ModuleGroup, error) {
	return r.apiClient.GetModuleGroupBatched(ctx, obj.ModuleGroupID)
}

// OrgAccount is the resolver for the orgAccount field.
func (r *moduleAssignmentResolver) OrgAccount(ctx context.Context, obj *models.ModuleAssignment) (*models.OrgAccount, error) {
	return r.apiClient.GetOrgAccountBatched(ctx, obj.OrgAccountID)
}

// ModulePropagation is the resolver for the modulePropagation field.
func (r *moduleAssignmentResolver) ModulePropagation(ctx context.Context, obj *models.ModuleAssignment) (*models.ModulePropagation, error) {
	if obj.ModulePropagationID == nil {
		return nil, nil
	}
	return r.apiClient.GetModulePropagationBatched(ctx, *obj.ModulePropagationID)
}

// TerraformDriftCheckRequests is the resolver for the terraformDriftCheckRequests field.
func (r *moduleAssignmentResolver) TerraformDriftCheckRequests(ctx context.Context, obj *models.ModuleAssignment, filters *models.TerraformDriftCheckRequestFilters, limit *int, offset *int) ([]*models.TerraformDriftCheckRequest, error) {
	return r.apiClient.GetTerraformDriftCheckRequestsForModuleAssignment(ctx, obj.ID, filters, limit, offset)
}

// TerraformExecutionRequests is the resolver for the terraformExecutionRequests field.
func (r *moduleAssignmentResolver) TerraformExecutionRequests(ctx context.Context, obj *models.ModuleAssignment, filters *models.TerraformExecutionRequestFilters, limit *int, offset *int) ([]*models.TerraformExecutionRequest, error) {
	return r.apiClient.GetTerraformExecutionRequestsForModuleAssignment(ctx, obj.ID, filters, limit, offset)
}

// TerraformConfiguration is the resolver for the terraformConfiguration field.
func (r *moduleAssignmentResolver) TerraformConfiguration(ctx context.Context, obj *models.ModuleAssignment) (string, error) {
	var modulePropagation *models.ModulePropagation
	var err error
	if obj.ModulePropagationID != nil {
		modulePropagation, err = r.apiClient.GetModulePropagation(ctx, *obj.ModulePropagationID)
		if err != nil {
			return "", err
		}
	}

	moduleVersion, err := r.apiClient.GetModuleVersion(ctx, obj.ModuleVersionID)
	if err != nil {
		return "", err
	}

	orgAccount, err := r.apiClient.GetOrgAccount(ctx, obj.OrgAccountID)
	if err != nil {
		return "", err
	}

	terraformConfig, err := terraform.GetTerraformConfiguration(&terraform.TerraformConfigurationInput{
		ModuleAssignment:  obj,
		ModulePropagation: modulePropagation,
		ModuleVersion:     moduleVersion,
		OrgAccount:        orgAccount,
		LockTableName:     r.config.Prefix + "-terraform-lock",
	})

	return string(terraformConfig), err
}

// StateVersions is the resolver for the stateVersions field.
func (r *moduleAssignmentResolver) StateVersions(ctx context.Context, obj *models.ModuleAssignment, limit *int, offset *int) ([]*models.StateVersion, error) {
	return r.apiClient.GetStateFileVersions(ctx, obj.RemoteStateBucket, obj.RemoteStateKey, limit)
}

// CreateModuleAssignment is the resolver for the createModuleAssignment field.
func (r *mutationResolver) CreateModuleAssignment(ctx context.Context, moduleAssignment models.NewModuleAssignment) (*models.ModuleAssignment, error) {
	return r.apiClient.CreateModuleAssignment(ctx, &moduleAssignment)
}

// UpdateModuleAssignment is the resolver for the updateModuleAssignment field.
func (r *mutationResolver) UpdateModuleAssignment(ctx context.Context, moduleAssignmentID uint, moduleAssignmentUpdate models.ModuleAssignmentUpdate) (*models.ModuleAssignment, error) {
	return r.apiClient.UpdateModuleAssignment(ctx, moduleAssignmentID, &moduleAssignmentUpdate)
}

// DeleteModuleAssignment is the resolver for the deleteModuleAssignment field.
func (r *mutationResolver) DeleteModuleAssignment(ctx context.Context, moduleAssignmentID uint) (bool, error) {
	err := r.apiClient.DeleteModuleAssignment(ctx, moduleAssignmentID)
	return err == nil, err
}

// ModuleAssignment is the resolver for the moduleAssignment field.
func (r *queryResolver) ModuleAssignment(ctx context.Context, moduleAssignmentID uint) (*models.ModuleAssignment, error) {
	return r.apiClient.GetModuleAssignment(ctx, moduleAssignmentID)
}

// ModuleAssignments is the resolver for the moduleAssignments field.
func (r *queryResolver) ModuleAssignments(ctx context.Context, filters *models.ModuleAssignmentFilters, limit *int, offset *int) ([]*models.ModuleAssignment, error) {
	return r.apiClient.GetModuleAssignments(ctx, filters, limit, offset)
}

// ModuleAssignment returns generated.ModuleAssignmentResolver implementation.
func (r *Resolver) ModuleAssignment() generated.ModuleAssignmentResolver {
	return &moduleAssignmentResolver{r}
}

type moduleAssignmentResolver struct{ *Resolver }
