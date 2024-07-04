package listener

// ConsumerInterface defines the interface for a consumer.
type ConsumerInterface interface {
	Consume()
	GetListenerTag() string
	GetMsgCh() <-chan []byte
}
