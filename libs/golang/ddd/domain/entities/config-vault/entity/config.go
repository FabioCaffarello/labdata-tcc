package entity

import (
	"errors"
	"fmt"
	regularTypesConversion "libs/golang/ddd/shared/type-tools/regular-types-converter/conversion"
	md5id "libs/golang/shared/id/go-md5"
	uuid "libs/golang/shared/id/go-uuid"
	typetools "libs/golang/shared/type-tools"
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
	ID        md5id.ID          `bson:"_id"`
	Active    bool              `bson:"active"`
	Service   string            `bson:"service"`
	Source    string            `bson:"source"`
	Provider  string            `bson:"provider"`
	DependsOn []JobDependencies `bson:"depends_on"`
	ConfigVersionID uuid.ID     `bson:"config_version_id"`
	CreatedAt       time.Time   `bson:"created_at"`
	UpdatedAt       time.Time   `bson:"updated_at"`
}

// ConfigProps represents the properties needed to create a new Config entity.
type ConfigProps struct {
	Active    bool
	Service   string
	Source    string
	Provider  string
	DependsOn []map[string]interface{}
	UpdatedAt string
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
	fmt.Printf("Config Props: %+v\n", configProps)

	idData := getIDData(configProps.Service, configProps.Source, configProps.Provider)
	updatedDate, err := typetools.ParseDateWithFormat(configProps.UpdatedAt, dateLayout)
	if err != nil {
		return nil, err
	}

	dependsOn, err := transformDependsOn(configProps.DependsOn)
	fmt.Printf("DependsOn: %+v\n", dependsOn)
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
		UpdatedAt: updatedDate,
		CreatedAt: time.Now(),
	}

	versionID, err := uuid.GenerateUUIDFromMap(config.GetVersionIDData())
	fmt.Printf("Version ID: %s\n", versionID)
	if err != nil {
		return nil, err
	}
	config.SetConfigVersionID(versionID)

	fmt.Printf("Config: %+v\n", config)

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

// ToMap converts the Config entity to a map representation.
func (c *Config) ToMap() (map[string]interface{}, error) {
	doc, err := regularTypesConversion.ConvertFromEntityToMapString(c)
	if err != nil {
		return nil, err
	}
	return doc, nil
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
	if c.CreatedAt.IsZero() {
		return ErrInvalidCreatedAt
	}
	return nil
}
