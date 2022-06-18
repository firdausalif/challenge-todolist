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

type TodoRepoitory struct {
	*database.DB
}

func NewTodoRepository(db *database.DB) models.TodoRepository {
	return &TodoRepoitory{DB: db}
}

func (t *TodoRepoitory) InsertTodos(ctx context.Context, a *models.Todo) (*models.Todo, error) {
	if err := t.DB.WithContext(ctx).Create(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (t *TodoRepoitory) GetTodoByID(ctx context.Context, id uint64) (*models.Todo, error) {
	todo := new(models.Todo)
	if err := t.DB.WithContext(ctx).
		Select("id", "activity_group_id", "title", "priority", "created_at", "updated_at", "deleted_at").
		Where("id = ?", id).
		First(todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewErrorRecordNotFound(constant.CodeErrDataNotFound, err)
		}
		return nil, helper.NewErrorMsg(constant.CodeErrQueryDB, err)
	}

	return todo, nil
}

func (t *TodoRepoitory) UpdateTodo(ctx context.Context, a *models.Todo) (*models.Todo, error) {
	res := t.DB.WithContext(ctx).
		Model(&models.Todo{ID: a.ID}).Updates(a).Find(a)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, helper.NewErrorRecordNotFound(constant.CodeErrDataNotFound, fmt.Errorf("0 rows effected"))
	}

	return a, nil
}

func (t *TodoRepoitory) RemoveTodo(ctx context.Context, id uint64) error {
	res := t.DB.WithContext(ctx).Delete(&models.Todo{ID: id})
	if res.Error != nil {
		return helper.NewErrorMsg(constant.CodeErrQueryDB, res.Error)
	}
	if res.RowsAffected == 0 {
		return helper.NewErrorRecordNotFound(constant.CodeErrDataNotFound, fmt.Errorf("0 rows effected"))
	}
	return nil
}

func (t *TodoRepoitory) FetchTodos(ctx context.Context, ActGroupID uint64) ([]*models.Todo, error) {
	var data []*models.Todo

	query := t.DB.WithContext(ctx).
		Select("id", "activity_group_id", "title", "priority", "created_at", "updated_at", "deleted_at")
	if ActGroupID != 0 {
		query = query.Where("activity_group_id = ? ", ActGroupID)
	}
	result := query.Find(&data)

	if result.Error != nil {
		return nil, helper.NewErrorMsg(constant.CodeErrQueryDB, result.Error)
	}

	return data, nil
}

func (t *TodoRepoitory) GetActivityGroup(ctx context.Context) ([]*models.Todo, error) {
	var data []*models.Todo

	if err := t.DB.WithContext(ctx).
		Select("activity_group_id").
		Group("activity_group_id").
		Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewErrorRecordNotFound(constant.CodeErrDataNotFound, err)
		}
		return nil, helper.NewErrorMsg(constant.CodeErrQueryDB, err)
	}

	return data, nil
}
