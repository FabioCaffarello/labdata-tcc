package actions

import (
	"libs/golang/clients/apis/config-vault/client"
	outputdto "libs/golang/ddd/dtos/config-vault/output"
)

type ListAllByDependenciesAction struct {
	client *client.Client
}

func NewListAllByDependenciesAction() *ListAllByDependenciesAction {
	return &ListAllByDependenciesAction{
		client: client.NewClient(),
	}
}

func (a *ListAllByDependenciesAction) Execute(provider, service, source string) ([]outputdto.ConfigDTO, error) {
	configs, err := a.client.ListConfigsByProviderAndDependencies(provider, service, source)
	if err != nil {
		return nil, err
	}
	return configs, nil
}
