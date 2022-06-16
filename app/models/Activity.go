package models

import (
	"context"
	"github.com/firdausalif/challenge-todolist/app/requests"
	"gorm.io/gorm"
	"time"
)

type (
	Activity struct {
		ID        int            `json:"id"`
		Email     string         `json:"email"`
		Title     string         `json:"title"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `json:"deleted_at"`
	}

	ActivityRepository interface {
		InsertActivity(ctx context.Context, a *Activity) (*Activity, error)
		GetActivityByID(ctx context.Context, id int) (*Activity, error)
		UpdateActivity(ctx context.Context, a *Activity) (*Activity, error)
		RemoveActivity(ctx context.Context, id int) error
		FetchActivity(ctx context.Context) ([]*Activity, error)
	}

	ActivityService interface {
		Create(ctx context.Context, req requests.CreateActivity) (*Activity, int, error)
		Update(ctx context.Context, id int, req requests.UpdateActivity) (*Activity, int, error)
		Remove(ctx context.Context, id int) (int, error)
		Fetch(ctx context.Context) ([]*Activity, int, error)
		GetById(ctx context.Context, id int) (*Activity, int, error)
	}
)
