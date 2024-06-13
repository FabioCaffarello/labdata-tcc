package entity

import (
	"errors"
	regularTypesConversion "libs/golang/ddd/shared/type-tools/regular-types-converter/conversion"
	md5id "libs/golang/shared/id/go-md5"
	uuid "libs/golang/shared/id/go-uuid"
	"reflect"
	"time"
)

var (
	// ErrInvalidID is returned when the ID of a Config is invalid.
	ErrInvalidID = errors.New("invalid ID")

	// ErrInvalidService is returned when the service of a Config is invalid.
	ErrInvalidService = errors.New("invalid service")

	// ErrInvalidSource is returned when the source of a Config is invalid.
	ErrInvalidSource = errors.New("invalid source")

	// ErrInvalidProvider is returned when the provider of a Config is invalid.
	ErrInvalidProvider = errors.New("invalid provider")

	// ErrInvalidConfigVersionID is returned when the config version ID of a Config is invalid.
	ErrInvalidConfigVersionID = errors.New("invalid config version ID")

	// ErrInvalidCreatedAt is returned when the created at timestamp of a Config is invalid.
	ErrInvalidCreatedAt = errors.New("invalid created at")

	// dateLayout defines the layout for parsing and formatting dates.
	dateLayout = "2006-01-02 15:04:05"
)

// JobDependencies represents the dependencies of a job, including the service and source.
type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}

// Config represents a configuration entity with various attributes such as service, source, provider, and dependencies.
type Config struct {
	ID              md5id.ID          `bson:"_id"`
	Active          bool              `bson:"active"`
	Service         string            `bson:"service"`
	Source          string            `bson:"source"`
	Provider        string            `bson:"provider"`
	DependsOn       []JobDependencies `bson:"depends_on"`
	ConfigVersionID uuid.ID           `bson:"config_version_id"`
	CreatedAt       string            `bson:"created_at"`
	UpdatedAt       string            `bson:"updated_at"`
}

// ConfigProps represents the properties needed to create a new Config entity.
type ConfigProps struct {
	Active    bool
	Service   string
	Source    string
	Provider  string
	DependsOn []map[string]interface{}
	// UpdatedAt string
}

// getIDData constructs a map with the service, source, and provider information.
func getIDData(service, source, provider string) map[string]string {
	return map[string]string{
		"service":  service,
		"source":   source,
		"provider": provider,
	}
}

// transformDependsOn converts a slice of map representations of dependencies to a slice of JobDependencies.
func transformDependsOn(dependsOn []map[string]interface{}) ([]JobDependencies, error) {
	dependsOnResult := make([]JobDependencies, len(dependsOn))
	for i, dep := range dependsOn {
		service, ok := dep["service"].(string)
		if !ok {
			return nil, errors.New("invalid service in depends_on")
		}
		source, ok := dep["source"].(string)
		if !ok {
			return nil, errors.New("invalid source in depends_on")
		}
		dependsOnResult[i] = JobDependencies{
			Service: service,
			Source:  source,
		}
	}
	return dependsOnResult, nil
}

// NewConfig creates a new Config entity based on the provided ConfigProps. It validates the
// properties and generates necessary IDs.
func NewConfig(configProps ConfigProps) (*Config, error) {
	idData := getIDData(configProps.Service, configProps.Source, configProps.Provider)

	dependsOn, err := transformDependsOn(configProps.DependsOn)
	if err != nil {
		return nil, err
	}

	config := &Config{
		ID:        md5id.NewID(idData),
		Active:    configProps.Active,
		Service:   configProps.Service,
		Source:    configProps.Source,
		Provider:  configProps.Provider,
		DependsOn: dependsOn,
		UpdatedAt: time.Now().Format(dateLayout),
		CreatedAt: time.Now().Format(dateLayout),
	}

	versionID, err := uuid.GenerateUUIDFromMap(config.GetVersionIDData())
	if err != nil {
		return nil, err
	}
	config.SetConfigVersionID(versionID)

	if err := config.isValid(); err != nil {
		return nil, err
	}

	return config, nil
}

// SetDependsOn sets the dependencies of the Config entity.
func (c *Config) SetDependsOn(dependsOn []JobDependencies) {
	c.DependsOn = dependsOn
}

// GetVersionIDData returns a map with the version ID data for the Config entity.
func (c *Config) GetVersionIDData() map[string]interface{} {
	return map[string]interface{}{
		"service":    c.Service,
		"source":     c.Source,
		"provider":   c.Provider,
		"active":     c.Active,
		"depends_on": c.DependsOn,
	}
}

// SetConfigVersionID sets the config version ID of the Config entity.
func (c *Config) SetConfigVersionID(configVersionID uuid.ID) {
	c.ConfigVersionID = configVersionID
}

// GetEntityID returns the ID of the Config entity.
func (c *Config) GetEntityID() string {
	return string(c.ID)
}

// SetCreatedAt sets the created at timestamp of the Config entity.
func (c *Config) SetCreatedAt(createdAt string) {
	c.CreatedAt = createdAt
}

// SetUpdatedAt sets the updated at timestamp of the Config entity.
func (c *Config) SetUpdatedAt(updatedAt string) {
	c.UpdatedAt = updatedAt
}

// ToMap converts the Config entity to a map representation.
func (c *Config) ToMap() (map[string]interface{}, error) {
	doc, err := regularTypesConversion.ConvertFromEntityToMapString(c)
	if err != nil {
		return nil, err
	}

	doc["_id"] = string(doc["_id"].(md5id.ID))
	doc["config_version_id"] = string(doc["config_version_id"].(uuid.ID))

	// Convert depends_on to a slice of maps
	dependsOn := make([]map[string]interface{}, len(c.DependsOn))
	for i, dep := range c.DependsOn {
		dependsOn[i] = map[string]interface{}{
			"service": dep.Service,
			"source":  dep.Source,
		}
	}
	doc["depends_on"] = dependsOn

	return doc, nil
}

// MapToEntity converts a map representation to a Config entity.
func (c *Config) MapToEntity(doc map[string]interface{}) (*Config, error) {
	if id, ok := doc["_id"].(string); ok {
		doc["_id"] = md5id.ID(id)
	} else {
		return nil, errors.New("field _id has invalid type")
	}

	if configVersionID, ok := doc["config_version_id"].(string); ok {
		doc["config_version_id"] = uuid.ID(configVersionID)
	} else {
		return nil, errors.New("field config_version_id has invalid type")
	}

	dependsOnSlice, ok := doc["depends_on"].([]interface{})
	if !ok {
		dependsOnMaps, ok := doc["depends_on"].([]map[string]interface{})
		if !ok {
			return nil, errors.New("field depends_on has invalid type")
		}
		dependsOnSlice = make([]interface{}, len(dependsOnMaps))
		for i, dep := range dependsOnMaps {
			dependsOnSlice[i] = dep
		}
	}

	jobDeps := make([]JobDependencies, len(dependsOnSlice))
	for i, dep := range dependsOnSlice {
		depMap, ok := dep.(map[string]interface{})
		if !ok {
			return nil, errors.New("field depends_on has invalid type or missing required fields")
		}
		service, serviceOK := depMap["service"].(string)
		source, sourceOK := depMap["source"].(string)
		if !serviceOK || !sourceOK {
			return nil, errors.New("field depends_on has invalid type or missing required fields")
		}
		jobDeps[i] = JobDependencies{
			Service: service,
			Source:  source,
		}
	}
	doc["depends_on"] = jobDeps

	configEntity, err := regularTypesConversion.ConvertFromMapStringToEntity(reflect.TypeOf(Config{}), doc)
	if err != nil {
		return nil, err
	}

	config := configEntity.(*Config)
	config.DependsOn = jobDeps

	return config, nil
}

// isValid validates the Config entity, ensuring all required fields are set.
func (c *Config) isValid() error {
	if c.ID == "" {
		return ErrInvalidID
	}
	if c.Service == "" {
		return ErrInvalidService
	}
	if c.Source == "" {
		return ErrInvalidSource
	}
	if c.Provider == "" {
		return ErrInvalidProvider
	}
	if c.ConfigVersionID == "" {
		return ErrInvalidConfigVersionID
	}
	return nil
}
