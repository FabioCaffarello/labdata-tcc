package event

import "time"

type OrderedProcess struct {
	Name    string
	Payload interface{}
}

func NewOrderedProcess() *OrderedProcess {
	return &OrderedProcess{
		Name: "OrderedProcess",
	}
}

func (e *OrderedProcess) GetName() string {
	return e.Name
}

func (e *OrderedProcess) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderedProcess) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderedProcess) GetDateTime() time.Time {
	return time.Now()
}
