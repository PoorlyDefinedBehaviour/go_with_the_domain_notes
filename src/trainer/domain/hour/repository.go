package hour

import (
	"context"
	"time"
)

type Repository interface {
	GetByTimes(ctx context.Context, times []time.Time) ([]Hour, error)
	GetByTime(ctx context.Context, time time.Time) (Hour, error)
	UpdateHour(ctx context.Context, hour Hour) error
}
