package pubsub

import (
	"context"
	"deathstar"
)

type Service interface {
	Subscribe(ctx context.Context) (*deathstar.Event, error)
}
