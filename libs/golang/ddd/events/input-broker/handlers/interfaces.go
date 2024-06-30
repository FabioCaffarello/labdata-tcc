package handler

// NotifierInterface defines the methods that a notifier should implement.
type NotifierInterface interface {
	Notify(message []byte, routingKey string) error
}
