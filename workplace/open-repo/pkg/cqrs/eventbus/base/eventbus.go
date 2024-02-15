package base

type EventBus interface {
	Publish(topic string, msg interface{}) error
	PublishRaw(topic string, msg []byte) error
	//
	// Subscribe(topic string, handler func(message []byte) error) error
	// Unsubscribe(topic string) error
	// //
	// Close() error
}
