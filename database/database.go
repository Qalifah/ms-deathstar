package database

import (
	"context"
	"deathstar"
)

type Repository interface {
	AddTargets(ctx context.Context, targets []deathstar.Target) error
}
