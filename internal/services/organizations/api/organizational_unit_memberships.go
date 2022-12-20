package api

import (
	"context"

	"github.com/sheacloud/tfom/pkg/organizations/models"
)

func (c *OrganizationsAPIClient) GetOrganizationalUnitMembershipsByAccount(ctx context.Context, accountId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error) {
	return c.dbClient.GetOrganizationalUnitMembershipsByAccount(ctx, accountId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetOrganizationalUnitMembershipsByOrgUnit(ctx context.Context, orgUnitId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error) {
	return c.dbClient.GetOrganizationalUnitMembershipsByOrgUnit(ctx, orgUnitId, limit, cursor)
}

func (c *OrganizationsAPIClient) GetOrganizationalUnitMembershipsByDimension(ctx context.Context, dimensionId string, limit int32, cursor string) (*models.OrganizationalUnitMemberships, error) {
	return c.dbClient.GetOrganizationalUnitMembershipsByDimension(ctx, dimensionId, limit, cursor)
}

func (c *OrganizationsAPIClient) PutOrganizationalUnitMembership(ctx context.Context, input *models.NewOrganizationalUnitMembership) (*models.OrganizationalUnitMembership, error) {
	orgUnitMembership := models.OrganizationalUnitMembership{
		OrgAccountId: input.OrgAccountId,
		DimensionId:  input.DimensionId,
		OrgUnitId:    input.OrgUnitId,
	}
	err := c.dbClient.PutOrganizationalUnitMembership(ctx, &orgUnitMembership)
	if err != nil {
		return nil, err
	} else {
		return &orgUnitMembership, nil
	}
}

func (c *OrganizationsAPIClient) DeleteOrganizationalUnitMembership(ctx context.Context, dimensionId string, accountId string) error {
	return c.dbClient.DeleteOrganizationalUnitMembership(ctx, dimensionId, accountId)
}