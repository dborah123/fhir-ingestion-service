package publisher

import "context"

type EventPublisher interface {
	Publish(ctx context.Context, event any) error
	Ping(ctx context.Context) error // For our health check!
}
