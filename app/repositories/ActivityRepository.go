package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/firdausalif/challenge-todolist/app/models"
	"github.com/firdausalif/challenge-todolist/pkg/constant"
	"github.com/firdausalif/challenge-todolist/pkg/helper"
	"github.com/firdausalif/challenge-todolist/platform/database"
	"gorm.io/gorm"
)

type ActivityRepoitory struct {
	*database.DB
}

func NewActivityRepository(db *database.DB) models.ActivityRepository {
	return &ActivityRepoitory{DB: db}
}

func (ac *ActivityRepoitory) InsertActivity(ctx context.Context, a *models.Activity) (*models.Activity, error) {
	if err := ac.DB.WithContext(ctx).Create(&a).Error; err != nil {
		return nil, helper.NewErrorMsg(constant.CodeErrQueryDB, err)
	}
	return a, nil
}

func (ac *ActivityRepoitory) GetActivityByID(ctx context.Context, id int) (*models.Activity, error) {
	act := new(models.Activity)
	if err := ac.DB.WithContext(ctx).
		Select("id", "email", "title", "created_at", "updated_at", "deleted_at").
		Where("id = ?", id).
		First(act).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewErrorRecordNotFound(constant.CodeErrDataNotFound, err)
		}
		return nil, helper.NewErrorMsg(constant.CodeErrQueryDB, err)
	}

	return act, nil
}

func (ac *ActivityRepoitory) UpdateActivity(ctx context.Context, a *models.Activity) (*models.Activity, error) {
	res := ac.DB.WithContext(ctx).
		Model(&models.Activity{ID: a.ID}).Updates(a).Find(a)

	if res.Error != nil {
		return nil, helper.NewErrorMsg(constant.CodeErrQueryDB, res.Error)
	}

	if res.RowsAffected == 0 {
		return nil, helper.NewErrorRecordNotFound(constant.CodeErrDataNotFound, fmt.Errorf("0 rows effected"))
	}

	return a, nil
}

func (ac *ActivityRepoitory) RemoveActivity(ctx context.Context, id int) error {
	res := ac.DB.WithContext(ctx).Delete(&models.Activity{ID: id})
	if res.Error != nil {
		return helper.NewErrorMsg(constant.CodeErrQueryDB, res.Error)
	}
	if res.RowsAffected == 0 {
		return helper.NewErrorMsg(constant.CodeErrDataNotFound, fmt.Errorf("0 rows effected"))
	}
	return nil
}

func (ac *ActivityRepoitory) FetchActivity(ctx context.Context) ([]*models.Activity, error) {
	var data []*models.Activity
	if err := ac.DB.WithContext(ctx).
		Select("id", "email", "title", "email", "created_at", "updated_at").
		Find(&data).Error; err != nil {
		return nil, helper.NewErrorMsg(constant.CodeErrQueryDB, err)
	}
	return data, nil
}
