package command

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/poorlydefinedbehaviour/wild_workouts/src/trainer/domain/hour"
)

type CancelTrainingHandler struct {
	hourRepo hour.Repository
}

func NewCancelTrainingHandler(hourRepo hour.Repository) CancelTrainingHandler {
	if hourRepo == nil {
		panic("nil hour.Repository")
	}

	return CancelTrainingHandler{hourRepo: hourRepo}
}

func (h *CancelTrainingHandler) Handle(ctx context.Context, hourToCancel time.Time) error {
	hour, err := h.hourRepo.GetByTime(ctx, hourToCancel)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := hour.CancelTraining(); err != nil {
		return errors.WithStack(err)
	}

	if err := h.hourRepo.UpdateHour(ctx, hour); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
