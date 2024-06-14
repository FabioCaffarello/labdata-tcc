package entity

type ConfigRepositoryInterface interface {
	Create(config *Config) error
	FindByID(id string) (*Config, error)
	FindAll() ([]*Config, error)
	Update(config *Config) error
	Delete(id string) error
	FindAllByServiceAndProvider(provider string, service string) ([]*Config, error)
	FindAllBySourceAndProvider(provider string, source string) ([]*Config, error)
	FindAllByServiceAndSourceAndProvider(service, source, provider string) ([]*Config, error)
	FindAllByServiceAndProviderAndActive(service, provider string, active bool) ([]*Config, error)
	FindAllByDependsOn(dependsOn map[string]interface{}) ([]*Config, error)
}
