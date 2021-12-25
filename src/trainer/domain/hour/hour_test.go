package hour_test

import (
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/poorlydefinedbehaviour/wild_workouts/src/trainer/domain/hour"
	"github.com/stretchr/testify/assert"
)

func Test_Hour_CancelTraining(t *testing.T) {
	t.Parallel()

	t.Run("traning cannot be canceled if there's no training scheduled for the hour", func(t *testing.T) {
		t.Parallel()

		h := hour.NewAvailable(time.Now())

		assert.True(t, errors.Is(h.CancelTraining(), hour.ErrNoTrainingScheduled))
	})

	t.Run("cancels training", func(t *testing.T) {
		t.Parallel()

		h := hour.NewAvailable(time.Now())

		assert.NoError(t, h.ScheduleTraining())
		assert.NoError(t, h.CancelTraining())

		assert.Equal(t, hour.Available, h.Availability())
	})
}

func Test_Hour_ScheduleTraining(t *testing.T) {
	t.Parallel()

	t.Run("training cannot be scheduled if hour already has a training scheduled", func(t *testing.T) {
		t.Parallel()

		h := hour.NewAvailable(time.Now())

		assert.NoError(t, h.ScheduleTraining())

		assert.True(t, errors.Is(h.ScheduleTraining(), hour.ErrHourHasTrainingScheduled))
	})

	t.Run("schedules training", func(t *testing.T) {
		t.Parallel()

		h := hour.NewAvailable(time.Now())

		assert.NoError(t, h.ScheduleTraining())

		assert.Equal(t, hour.TrainingScheduled, h.Availability())
	})
}

func Test_Hour_MakeAvailable(t *testing.T) {
	t.Parallel()

	t.Run("hour cannot be made available when there's a training scheduled for it", func(t *testing.T) {
		t.Parallel()

		h := hour.NewAvailable(time.Now())

		assert.NoError(t, h.ScheduleTraining())

		err := h.MakeAvailable()

		assert.True(t, errors.Is(err, hour.ErrHourHasTrainingScheduled))
	})

	t.Run("makes hour available", func(t *testing.T) {
		t.Parallel()

		h := hour.NewAvailable(time.Now())

		assert.NoError(t, h.MakeUnavailable())

		assert.NoError(t, h.MakeAvailable())

		assert.Equal(t, hour.Available, h.Availability())
	})
}

func Test_Hour_MakeUnavailable(t *testing.T) {
	t.Parallel()

	t.Run("hour cannot be made unavailable when there's a training scheduled for it", func(t *testing.T) {
		t.Parallel()

		h := hour.NewAvailable(time.Now())

		assert.NoError(t, h.ScheduleTraining())

		err := h.MakeUnavailable()

		assert.True(t, errors.Is(err, hour.ErrHourHasTrainingScheduled))
	})

	t.Run("makes hour unavailable", func(t *testing.T) {
		t.Parallel()

		h := hour.NewAvailable(time.Now())

		assert.Equal(t, hour.Available, h.Availability())

		assert.NoError(t, h.MakeUnavailable())

		assert.Equal(t, hour.Unavailable, h.Availability())
	})
}
