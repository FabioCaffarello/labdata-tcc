package usecaseprotocol

// UseCaseProtocol defines the interface for a use case protocol.
type UseCaseProtocol interface {
	ProcessMessageChannel(msgCh <-chan []byte, listenerTag string)
}
