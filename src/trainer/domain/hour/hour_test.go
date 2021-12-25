package hour_test

import (
	"errors"
	"testing"
	"time"

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

		assert.True(t, errors.Is(h.ScheduleTraining(), hour.ErrTrainingAlreadyScheduledForHour))
	})

	t.Run("schedules training", func(t *testing.T) {
		t.Parallel()

		h := hour.NewAvailable(time.Now())

		assert.NoError(t, h.ScheduleTraining())

		assert.Equal(t, hour.TrainingScheduled, h.Availability())
	})
}
