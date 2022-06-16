package services

import (
	"context"
	"github.com/firdausalif/challenge-todolist/app/models"
	"github.com/firdausalif/challenge-todolist/app/requests"
	"github.com/firdausalif/challenge-todolist/pkg/constant"
	"time"
)

type ActivityService struct {
	activityRepository models.ActivityRepository
	ctxTimeout         time.Duration
}

func NewActivityService(ar models.ActivityRepository, contextTimeout time.Duration) models.ActivityService {
	return &ActivityService{
		activityRepository: ar,
		ctxTimeout:         contextTimeout,
	}
}

func (as *ActivityService) Create(ctx context.Context, req requests.CreateActivity) (*models.Activity, int, error) {
	ctx, cancel := context.WithTimeout(ctx, as.ctxTimeout)
	defer cancel()

	gatra, err := as.activityRepository.InsertActivity(ctx, &models.Activity{
		Email: req.Email,
		Title: req.Title,
	})

	if err != nil {
		return nil, constant.CodeInternalServerError, err
	}

	return gatra, constant.CodeSuccess, nil
}

func (as *ActivityService) Update(ctx context.Context, id int, req requests.UpdateActivity) (*models.Activity, int, error) {
	ctx, cancel := context.WithTimeout(ctx, as.ctxTimeout)
	defer cancel()

	updateActivity, err := as.activityRepository.UpdateActivity(ctx, &models.Activity{
		ID:    id,
		Title: req.Title,
		Email: req.Email,
	})

	if err != nil {
		return nil, constant.CodeInternalServerError, err
	}

	return updateActivity, constant.CodeSuccess, nil
}

func (as *ActivityService) Remove(ctx context.Context, id int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, as.ctxTimeout)
	defer cancel()

	err := as.activityRepository.RemoveActivity(ctx, id)
	if err != nil {
		return constant.CodeInternalServerError, err
	}

	return constant.CodeSuccess, nil
}

func (as *ActivityService) Fetch(ctx context.Context) ([]*models.Activity, int, error) {
	ctx, cancel := context.WithTimeout(ctx, as.ctxTimeout)
	defer cancel()

	fetchActivity, err := as.activityRepository.FetchActivity(ctx)
	if err != nil {
		return nil, constant.CodeInternalServerError, err
	}

	return fetchActivity, constant.CodeSuccess, nil
}

func (as *ActivityService) GetById(ctx context.Context, id int) (*models.Activity, int, error) {
	ctx, cancel := context.WithTimeout(ctx, as.ctxTimeout)
	defer cancel()

	activity, err := as.activityRepository.GetActivityByID(ctx, id)
	if err != nil {
		return nil, constant.CodeInternalServerError, err
	}

	return activity, constant.CodeSuccess, nil
}
