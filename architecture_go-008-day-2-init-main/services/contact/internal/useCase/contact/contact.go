package contact

import (
	"github.com/google/uuid"

	"architecture_go/pkg/type/queryParameter"
	"architecture_go/services/contact/internal/domain/contact"
)

func (uc *UseCase) Create(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) Update(contactUpdate contact.Contact) (*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) Delete(ID uuid.UUID) error {
	panic("implement me")
}

func (uc *UseCase) List(parameter queryParameter.QueryParameter) ([]*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) ReadByID(ID uuid.UUID) (response *contact.Contact, err error) {

	panic("implement me")
}

func (uc *UseCase) Count() (uint64, error) {
	panic("implement me")
}
