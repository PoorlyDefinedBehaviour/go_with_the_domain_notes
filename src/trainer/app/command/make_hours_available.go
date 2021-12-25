package command

import (
	"context"
	"time"

	"github.com/poorlydefinedbehaviour/wild_workouts/src/trainer/domain/hour"
)

type MakeHoursAvailableHandler struct {
	hourRepo hour.Repository
}

func NewMakeHoursAvailableHandler(hourRepo hour.Repository) MakeHoursAvailableHandler {
	if hourRepo == nil {
		panic("nil hour.Repository")
	}

	return MakeHoursAvailableHandler{hourRepo: hourRepo}
}

func (h *MakeHoursAvailableHandler) Handle(ctx context.Context, hoursToMakeAvailable []time.Time) error {
	panic("TODO")
}
