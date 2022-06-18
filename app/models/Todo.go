package models

import (
	"context"
	"github.com/firdausalif/challenge-todolist/app/requests"
	"gorm.io/gorm"
	"time"
)

type (
	Todo struct {
		ID              uint64         `db:"id" json:"id"`
		ActivityGroupID uint64         `db:"activity_group_id" json:"activity_group_id" gorm:"index"`
		Title           string         `db:"title" json:"title"`
		IsActive        *bool          `db:"is_active" gorm:"default:1" json:"is_active"`
		Priority        string         `db:"priority" gorm:"default:'very-high'" json:"priority"`
		CreatedAt       time.Time      `db:"created_at" json:"created_at"`
		UpdatedAt       time.Time      `db:"updated_at" json:"updated_at"`
		DeletedAt       gorm.DeletedAt `db:"deleted_at" json:"deleted_at"`
	}

	TodoRepository interface {
		InsertTodos(ctx context.Context, a *Todo) (*Todo, error)
		GetTodoByID(ctx context.Context, id uint64) (*Todo, error)
		UpdateTodo(ctx context.Context, a *Todo) (*Todo, error)
		RemoveTodo(ctx context.Context, id uint64) error
		FetchTodos(ctx context.Context, ActGroupID uint64) ([]*Todo, error)
		GetActivityGroup(ctx context.Context) ([]*Todo, error)
	}

	TodoService interface {
		Create(ctx context.Context, req requests.CreateTodo) (*Todo, int, error)
		Update(ctx context.Context, id uint64, req requests.UpdateTodo) (*Todo, int, error)
		Remove(ctx context.Context, id uint64) (int, error)
		Fetch(ctx context.Context, id uint64) ([]*Todo, int, error)
		GetById(ctx context.Context, id uint64) (*Todo, int, error)
		GetActivityGroup(ctx context.Context) ([]*Todo, error)
	}
)
