package repo

import (
	"context"
	"time"
)

type Setter interface {
	Set(context.Context, string, interface{}, time.Time) error
}

type Getter interface {
	Get(context.Context, string) (interface{}, error)
}

type Deleter interface {
	Delete(context.Context, string) error
}

// Cleaner is used to Swipe up the expired entities
type Cleaner interface {
	Clean(context.Context) error
}

// Snapper take a snap from map when is called
// recieves Path and context
type Shotter interface {
	Shot(context.Context, string) error
}
