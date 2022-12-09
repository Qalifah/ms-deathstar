package pulsar

import (
	"context"
	"deathstar"
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
)

type Pulsar struct {
	consumer pulsar.Consumer
}

func (p *Pulsar) Subscribe(ctx context.Context) (*deathstar.Event, error) {
	msg, err := p.consumer.Receive(ctx)
	if err != nil {
		return nil, err
	}
	var event deathstar.Event
	if err := json.Unmarshal(msg.Payload(), &event); err != nil {
		return nil, err
	}
	p.consumer.Ack(msg)
	return &event, nil
}

func (p *Pulsar) Close() {
	p.consumer.Close()
}

func New(url, topic string) (*Pulsar, error) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: url,
	})
	if err != nil {
		return nil, err
	}

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: "my-sub",
		Type:             pulsar.Shared,
	})
	if err != nil {
		return nil, err
	}

	return &Pulsar{
		consumer: consumer,
	}, nil
}
