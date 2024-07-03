package entity

type EventOrderRepositoryInterface interface {
	Create(output *EventOrder) error
	FindByID(id string) (*EventOrder, error)
	FindAll() ([]*EventOrder, error)
	Delete(id string) error
}
