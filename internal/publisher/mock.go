package publisher

import (
	"context"
	"fmt"
)

type MockPublisher struct{}

func NewMockPublisher() *MockPublisher {
	return &MockPublisher{}
}

func (m *MockPublisher) Publish(ctx context.Context, event any) error {
	fmt.Printf("Mock published event: %v\n", event)
	return nil
}

func (m *MockPublisher) Ping(ctx context.Context) error {
	return nil // Mock is always healthy
}
