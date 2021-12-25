package command

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/poorlydefinedbehaviour/wild_workouts/src/trainer/domain/hour"
)

type ScheduleTrainingHandler struct {
	hourRepo hour.Repository
}

func NewScheduleTrainingHandler(hourRepo hour.Repository) ScheduleTrainingHandler {
	if hourRepo == nil {
		panic("nil hour.Repository")
	}

	return ScheduleTrainingHandler{hourRepo: hourRepo}
}

func (h *ScheduleTrainingHandler) Handle(ctx context.Context, hourToSchedule time.Time) error {
	hour, err := h.hourRepo.GetByTime(ctx, hourToSchedule)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := hour.ScheduleTraining(); err != nil {
		return errors.WithStack(err)
	}

	if err := h.hourRepo.UpdateHour(ctx, hour); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
