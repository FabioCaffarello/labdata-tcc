package entity

type ConfigRepositoryInterface interface {
	Create(config *Config) error
	FindByID(id string) (*Config, error)
	FindAll() ([]*Config, error)
	Update(config *Config) error
	Delete(id string) error
	FindAllByServiceAndProvider(provider, service string) ([]*Config, error)
	FindAllBySourceAndProvider(provider, source string) ([]*Config, error)
	FindAllByServiceAndSourceAndProvider(service, source, provider string) ([]*Config, error)
	FindAllByServiceAndProviderAndActive(service, provider string, active bool) ([]*Config, error)
	FindAllByProviderAndDependsOn(provider, service, source string) ([]*Config, error)
}
