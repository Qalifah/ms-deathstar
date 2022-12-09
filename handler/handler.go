package handler

import (
	"context"
	"deathstar/database"
	"deathstar/lib/pubsub"
	"deathstar/proto"
)

type handler struct {
	pubSubService pubsub.Service
	store         database.Repository
	proto.UnimplementedDeathstarServer
}

func (h *handler) Listen(ctx context.Context) error {
	event, err := h.pubSubService.Subscribe(ctx)
	if err != nil {
		return err
	}
	return h.store.AddTargets(ctx, event.Targets)
}

func New(targetStore database.Repository, psService pubsub.Service) *handler {
	return &handler{
		pubSubService: psService,
		store:         targetStore,
	}
}
