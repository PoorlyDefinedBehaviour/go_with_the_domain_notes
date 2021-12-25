package hour

import (
	"time"

	"github.com/pkg/errors"
)

type Availability struct{ value string }

func (availability Availability) String() string {
	return availability.value
}

var (
	Available         = Availability{"available"}
	Unavailable       = Availability{"unavailable"}
	TrainingScheduled = Availability{"training_scheduled"}
)

type Hour struct {
	hour         time.Time
	availability Availability
}

func NewAvailable(hour time.Time) Hour {
	return Hour{hour: hour, availability: Available}
}

func (hour Hour) Availability() Availability {
	return hour.availability
}

func (hour Hour) HasTrainingScheduled() bool {
	return hour.availability == TrainingScheduled
}

var ErrNoTrainingScheduled = errors.New("no training scheduled")

func (hour *Hour) CancelTraining() error {
	if !hour.HasTrainingScheduled() {
		return errors.WithStack(ErrNoTrainingScheduled)
	}

	hour.availability = Available

	return nil
}

var ErrTrainingAlreadyScheduledForHour = errors.New("hour already has a training scheduled")

func (hour *Hour) ScheduleTraining() error {
	if hour.HasTrainingScheduled() {
		return errors.WithStack(ErrTrainingAlreadyScheduledForHour)
	}

	hour.availability = TrainingScheduled

	return nil
}
