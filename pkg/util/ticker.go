package util

import (
	"context"
	"time"
)

func NewPeriodRoutine(ctx context.Context, period time.Duration, worker func()) {
	go func() {
		tick := time.NewTicker(period)
		for {
			select {
			case <-tick.C:
				worker()
			case <-ctx.Done():
				return
			}
		}
	}()
}
