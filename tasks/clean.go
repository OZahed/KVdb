package tasks

import (
	"KVdb/repo"
	"context"
	"time"
)

func CleanTask(ctx context.Context, d time.Duration, rc repo.Cleaner) func() error {
	return func() error {
		ctx, can := context.WithCancel(ctx)
		defer can()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(d):
			return rc.Clean(ctx)
		}
	}
}
