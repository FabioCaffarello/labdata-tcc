package entity

type SchemaRepositoryInterface interface {
	Create(schema *Schema) error
	FindByID(id string) (*Schema, error)
	FindAll() ([]*Schema, error)
	Update(schema *Schema) error
	Delete(id string) error
	FindAllByServiceAndProvider(provider, service string) ([]*Schema, error)
	FindAllBySourceAndProvider(provider, source string) ([]*Schema, error)
	FindAllByServiceAndSourceAndProvider(service, source, provider string) ([]*Schema, error)
	FindOneByServiceAndSourceAndProviderAndSchemaType(provider, service, source, schemaType string) (*Schema, error)
}
