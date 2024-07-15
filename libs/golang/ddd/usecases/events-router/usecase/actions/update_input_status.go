package actions

import (
	"libs/golang/clients/apis/input-broker/client"
	outputdto "libs/golang/ddd/dtos/events-router/output"
	shareddto "libs/golang/ddd/dtos/input-broker/shared"
)

type UpdateInputStatusAction struct {
	client *client.Client
}

func NewUpdateInputStatusAction() *UpdateInputStatusAction {
	return &UpdateInputStatusAction{
		client: client.NewClient(),
	}
}

func (a *UpdateInputStatusAction) Execute(inputMsg outputdto.ProcessOrderDTO, statusCode int, statusDetail string) error {
	status := shareddto.StatusDTO{
		Code:   statusCode,
		Detail: statusDetail,
	}
	_, err := a.client.UpdateInputStatus(inputMsg.InputID, status)
	if err != nil {
		return err
	}
	return nil
}
