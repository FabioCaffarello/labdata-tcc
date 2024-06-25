package entity

type OutputRepositoryInterface interface {
	Create(output *Output) error
	FindByID(id string) (*Output, error)
	FindAll() ([]*Output, error)
	Update(output *Output) error
	Delete(id string) error
	FindAllByServiceAndProvider(provider, service string) ([]*Output, error)
	FindAllBySourceAndProvider(provider, source string) ([]*Output, error)
	FindAllByServiceAndSourceAndProvider(service, source, provider string) ([]*Output, error)
}
