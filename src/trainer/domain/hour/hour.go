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

var (
	ErrHourHasTrainingScheduled = errors.New("hour has a training scheduled")
	ErrNoTrainingScheduled      = errors.New("no training scheduled")
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

func (hour *Hour) CancelTraining() error {
	if !hour.HasTrainingScheduled() {
		return errors.WithStack(ErrNoTrainingScheduled)
	}

	hour.availability = Available

	return nil
}

func (hour *Hour) ScheduleTraining() error {
	if hour.HasTrainingScheduled() {
		return errors.WithStack(ErrHourHasTrainingScheduled)
	}

	hour.availability = TrainingScheduled

	return nil
}

func (hour *Hour) MakeAvailable() error {
	if hour.HasTrainingScheduled() {
		return errors.WithStack(ErrHourHasTrainingScheduled)
	}

	hour.availability = Available

	return nil
}

func (hour *Hour) MakeUnavailable() error {
	if hour.HasTrainingScheduled() {
		return errors.WithStack(ErrHourHasTrainingScheduled)
	}

	hour.availability = Unavailable

	return nil
}
