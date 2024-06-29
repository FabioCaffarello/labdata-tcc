package entity

type InputRepositoryInterface interface {
	Create(output *Input) error
	FindByID(id string) (*Input, error)
	FindAll() ([]*Input, error)
	Update(output *Input) error
	Delete(id string) error
	FindByStatus(status int) ([]*Input, error)
	FindAllByServiceAndProvider(provider, service string) ([]*Input, error)
	FindAllBySourceAndProvider(provider, source string) ([]*Input, error)
	FindAllByServiceAndSourceAndProvider(provider, service, source string) ([]*Input, error)
	FindAllByStatusAndServiceAndProvider(service, provider string, status int) ([]*Input, error)
	FindAllByStatusAndSourceAndProvider(source, provider string, status int) ([]*Input, error)
	FindAllByStatusAndServiceAndSourceAndProvider(service, source, provider string, status int) ([]*Input, error)
}
