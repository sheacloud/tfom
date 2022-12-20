package api

import (
	"context"

	"github.com/sheacloud/tfom/internal/services/organizations/database"
	"github.com/sheacloud/tfom/pkg/organizations/models"
)

type OrganizationsAPIClientInterface interface {
	GetOrganizationalDimension(ctx context.Context, dimensionId string) (*models.OrganizationalDimension, error)
	GetOrganizationalDimensions(ctx context.Context, limit int32, cursor string) (*models.OrganizationalDimensions, error)
	PutOrganizationalDimension(ctx context.Context, input *models.NewOrganizationalDimension) (*models.OrganizationalDimension, error)
	DeleteOrganizationalDimension(ctx context.Context, dimensionId string) error

	GetOrganizationalUnit(ctx context.Context, orgUnitId string) (*models.OrganizationalUnit, error)
	GetOrganizationalUnits(ctx context.Context, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByDimension(ctx context.Context, dimensionId string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByParent(ctx context.Context, dimensionId string, parentOrgUnitId string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	GetOrganizationalUnitsByHierarchy(ctx context.Context, dimensionId string, hierarchy string, limit int32, cursor string) (*models.OrganizationalUnits, error)
	PutOrganizationalUnit(ctx context.Context, input *models.NewOrganizationalUnit) (*models.OrganizationalUnit, error)
	DeleteOrganizationalUnit(ctx context.Context, dimensionId string, orgUnitId string) error
	UpdateOrganizationalUnit(ctx context.Context, orgUnitId string, update *models.OrganizationalUnitUpdate) (*models.OrganizationalUnit, error)
	UpdateOrganizationalUnitHierarchies(ctx context.Context, dimensionId string) error

	GetOrganizationalAccount(ctx context.Context, orgAccountId string) (*models.OrganizationalAccount, error)
	GetOrganizationalAccounts(ctx context.Context, limit int32, cursor string) (*models.OrganizationalAccounts, error)
	PutOrganizationalAccount(ctx context.Context, input *models.NewOrganizationalAccount) (*models.OrganizationalAccount, error)
	DeleteOrganizationalAccount(ctx context.Context, orgAccountId string) error

	GetOrganizationalUnitMembershipsByAccount(ctx context.Context, accountId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error)
	GetOrganizationalUnitMembershipsByOrgUnit(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error)
	GetOrganizationalUnitMembershipsByDimension(ctx context.Context, dimensionId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error)
	PutOrganizationalUnitMembership(ctx context.Context, input *models.NewOrganizationalUnitMembership) (*models.OrganizationalUnitMembership, error)
	DeleteOrganizationalUnitMembership(ctx context.Context, dimensionId string, accountId string) error
}

type OrganizationsAPIClient struct {
	dbClient database.OrganizationsDatabaseClientInterface
}

func NewOrganizationsAPIClient(dbClient database.OrganizationsDatabaseClientInterface) *OrganizationsAPIClient {
	return &OrganizationsAPIClient{
		dbClient: dbClient,
	}
}