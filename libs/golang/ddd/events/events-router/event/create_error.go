package event

import "time"

type ErrorCreated struct {
	Name    string
	Payload interface{}
}

func NewErrorCreated() *ErrorCreated {
	return &ErrorCreated{
		Name: "ErrorCreated",
	}
}

func (e *ErrorCreated) GetName() string {
	return e.Name
}

func (e *ErrorCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *ErrorCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *ErrorCreated) GetDateTime() time.Time {
	return time.Now()
}
