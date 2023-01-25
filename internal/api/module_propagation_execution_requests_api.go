package api

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/sheacloud/tfom/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func applyModulePropagationExecutionRequestFilters(tx *gorm.DB, filters *models.ModulePropagationExecutionRequestFilters) *gorm.DB {
	if filters != nil {
		if filters.StartedBefore != nil {
			tx = tx.Where("started_at < ?", *filters.StartedBefore)
		}
		if filters.StartedAfter != nil {
			tx = tx.Where("started_at > ?", *filters.StartedAfter)
		}
		if filters.CompletedBefore != nil {
			tx = tx.Where("completed_at < ?", *filters.CompletedBefore)
		}
		if filters.CompletedAfter != nil {
			tx = tx.Where("completed_at > ?", *filters.CompletedAfter)
		}
		if filters.Status != nil {
			tx = tx.Where("status = ?", *filters.Status)
		}
	}
	return tx
}

func applyModulePropagationExecutionRequestPreloads(tx *gorm.DB) *gorm.DB {
	return tx
}

func (c *APIClient) GetModulePropagationExecutionRequestsByIds(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	output := make([]*dataloader.Result, len(keys))

	var modulePropagationExecutionRequests []*models.ModulePropagationExecutionRequest
	tx := applyModulePropagationExecutionRequestPreloads(c.db)
	err := tx.Find(&modulePropagationExecutionRequests, keys.Keys()).Error
	if err != nil {
		for i := range keys {
			output[i] = &dataloader.Result{Error: err}
		}
		return output
	}

	for i := range keys {
		output[i] = &dataloader.Result{Data: modulePropagationExecutionRequests[i], Error: nil}
	}
	return output
}

func (c *APIClient) GetModulePropagationExecutionRequest(ctx context.Context, id uint) (*models.ModulePropagationExecutionRequest, error) {
	var modulePropagationExecutionRequest models.ModulePropagationExecutionRequest
	tx := applyModulePropagationExecutionRequestPreloads(c.db)
	err := tx.First(&modulePropagationExecutionRequest, id).Error
	if err != nil {
		return nil, err
	}
	return &modulePropagationExecutionRequest, nil
}

func (c *APIClient) GetModulePropagationExecutionRequestBatched(ctx context.Context, id uint) (*models.ModulePropagationExecutionRequest, error) {
	thunk := c.modulePropagationExecutionRequestsLoader.Load(ctx, dataloader.StringKey(idToString(id)))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*models.ModulePropagationExecutionRequest), nil
}

func (c *APIClient) GetModulePropagationExecutionRequests(ctx context.Context, filters *models.ModulePropagationExecutionRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationExecutionRequest, error) {
	var modulePropagationExecutionRequests []*models.ModulePropagationExecutionRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyModulePropagationExecutionRequestFilters(tx, filters)
	tx = applyModulePropagationExecutionRequestPreloads(tx)
	err := tx.Find(&modulePropagationExecutionRequests).Error
	if err != nil {
		return nil, err
	}
	return modulePropagationExecutionRequests, nil
}

func (c *APIClient) GetModulePropagationExecutionRequestsForModulePropagation(ctx context.Context, modulePropagationId uint, filters *models.ModulePropagationExecutionRequestFilters, limit *int, offset *int) ([]*models.ModulePropagationExecutionRequest, error) {
	var modulePropagationExecutionRequests []*models.ModulePropagationExecutionRequest
	tx := applyPagination(c.db, limit, offset)
	tx = applyModulePropagationExecutionRequestFilters(tx, filters)
	tx = applyModulePropagationExecutionRequestPreloads(tx)
	err := tx.Model(&models.ModulePropagation{Model: gorm.Model{ID: modulePropagationId}}).Association("ModulePropagationExecutionRequestsAssociation").Find(&modulePropagationExecutionRequests)
	if err != nil {
		return nil, err
	}
	return modulePropagationExecutionRequests, nil
}

func (c *APIClient) CreateModulePropagationExecutionRequest(ctx context.Context, input *models.NewModulePropagationExecutionRequest) (*models.ModulePropagationExecutionRequest, error) {
	modulePropagationExecutionRequest := models.ModulePropagationExecutionRequest{
		ModulePropagationID: input.ModulePropagationID,
		Status:              models.RequestStatusPending,
	}
	err := c.db.Create(&modulePropagationExecutionRequest).Error
	if err != nil {
		return nil, err
	}

	return &modulePropagationExecutionRequest, nil
}

func (c *APIClient) DeleteModulePropagationExecutionRequest(ctx context.Context, id uint) error {
	return c.db.Select(clause.Associations).Delete(&models.ModulePropagationExecutionRequest{}, id).Error
}

func (c *APIClient) UpdateModulePropagationExecutionRequest(ctx context.Context, id uint, update *models.ModulePropagationExecutionRequestUpdate) (*models.ModulePropagationExecutionRequest, error) {
	modulePropagationExecutionRequest := models.ModulePropagationExecutionRequest{
		Model: gorm.Model{
			ID: id,
		},
	}
	updates := models.ModulePropagationExecutionRequest{}

	if update.Status != nil {
		updates.Status = *update.Status
	}
	if update.StartedAt != nil {
		updates.StartedAt = update.StartedAt
	}
	if update.CompletedAt != nil {
		updates.CompletedAt = update.CompletedAt
	}

	err := c.db.Model(&modulePropagationExecutionRequest).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &modulePropagationExecutionRequest, nil
}