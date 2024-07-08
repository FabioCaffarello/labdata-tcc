package actions

import (
	"libs/golang/clients/apis/schema-vault/client"
	outputdto "libs/golang/ddd/dtos/events-router/output"
	inputdto "libs/golang/ddd/dtos/schema-vault/input"
)

type ValidateSchemaAction struct {
	client *client.Client
}

func NewValidateSchemaAction() *ValidateSchemaAction {
	return &ValidateSchemaAction{
		client: client.NewClient(),
	}
}

func (a *ValidateSchemaAction) Execute(inputMsg outputdto.ProcessOrderDTO, schemaType string) error {
	schemaData := inputdto.SchemaDataDTO{
		Service:    inputMsg.Service,
		Source:     inputMsg.Source,
		Provider:   inputMsg.Provider,
		SchemaType: schemaType,
		Data:       inputMsg.Data,
	}
	err := a.client.ValidateSchema(schemaData)
	if err != nil {
		return err
	}
	return nil
}
