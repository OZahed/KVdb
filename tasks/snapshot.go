package tasks

import (
	"KVdb/repo"
	"context"
	"time"
)

func SnapShotTask(ctx context.Context, d time.Duration, rs repo.Shotter, path string) func() error {
	return func() error {
		ctx, can := context.WithCancel(ctx)
		defer can()
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(d):
			return rs.Shot(ctx, path)
		}
	}
}
