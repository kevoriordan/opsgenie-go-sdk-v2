package contact

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRequest_Validate(t *testing.T) {
	var err error
	createRequest := &CreateRequest{}
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("User identifier cannot be empty.").Error())

	createRequest.UserIdentifier = "123"
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("to cannot be empty.").Error())

	createRequest.To = "ab@email.com"
	err = createRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Method cannot be empty.").Error())

	createRequest.MethodOfContact = Email
	err = createRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, createRequest.Endpoint(), "/v2/users/123/contacts")
	assert.Equal(t, createRequest.Method(), "POST")

}

func TestGetRequest_Validate(t *testing.T) {
	var err error
	getRequest := &GetRequest{}
	err = GetRequest{}.Validate()
	assert.Equal(t, err.Error(), errors.New("User identifier cannot be empty.").Error())

	getRequest.UserIdentifier = "123"
	err = getRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Contact identifier cannot be empty.").Error())

	getRequest.ContactIdentifier = "1234"
	err = getRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, getRequest.Endpoint(), "/v2/users/123/contacts/1234")
	assert.Equal(t, getRequest.Method(), "GET")

}

func TestUpdateRequest_Validate(t *testing.T) {
	var err error
	updateRequest := &UpdateRequest{}
	err = GetRequest{}.Validate()
	assert.Equal(t, err.Error(), errors.New("User identifier cannot be empty.").Error())

	updateRequest.UserIdentifier = "123"
	err = updateRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Contact identifier cannot be empty.").Error())

	updateRequest.ContactIdentifier = "1234"
	err = updateRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("to cannot be empty.").Error())

	updateRequest.To = "ab@email.com"
	err = updateRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, updateRequest.Endpoint(), "/v2/users/123/contacts/1234")
	assert.Equal(t, updateRequest.Method(), "PATCH")

}
func TestDeleteRequest_Validate(t *testing.T) {
	var err error
	deleteRequest := &DeleteRequest{}
	err = DeleteRequest{}.Validate()
	assert.Equal(t, err.Error(), errors.New("User identifier cannot be empty.").Error())

	deleteRequest.UserIdentifier = "123"
	err = deleteRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Contact identifier cannot be empty.").Error())

	deleteRequest.ContactIdentifier = "1234"
	err = deleteRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, deleteRequest.Endpoint(), "/v2/users/123/contacts/1234")
	assert.Equal(t, deleteRequest.Method(), "DELETE")

}

func TestListRequest_Validate(t *testing.T) {
	var err error
	listRequest := &ListRequest{}
	err = ListRequest{}.Validate()
	assert.Equal(t, err.Error(), errors.New("User identifier cannot be empty.").Error())

	listRequest.UserIdentifier = "123"
	err = listRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, listRequest.Endpoint(), "/v2/users/123/contacts")
	assert.Equal(t, listRequest.Method(), "GET")

}

func TestEnableRequest_Validate(t *testing.T) {
	var err error
	enableRequest := &EnableRequest{}
	err = EnableRequest{}.Validate()
	assert.Equal(t, err.Error(), errors.New("User identifier cannot be empty.").Error())

	enableRequest.UserIdentifier = "123"
	err = enableRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Contact identifier cannot be empty.").Error())

	enableRequest.ContactIdentifier = "1234"
	err = enableRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, enableRequest.Endpoint(), "/v2/users/123/contacts/1234/enable")
	assert.Equal(t, enableRequest.Method(), "POST")

}

func TestDisableRequest_Validate(t *testing.T) {
	var err error
	enableRequest := &DisableRequest{}
	err = DisableRequest{}.Validate()
	assert.Equal(t, err.Error(), errors.New("User identifier cannot be empty.").Error())

	enableRequest.UserIdentifier = "123"
	err = enableRequest.Validate()
	assert.Equal(t, err.Error(), errors.New("Contact identifier cannot be empty.").Error())

	enableRequest.ContactIdentifier = "1234"
	err = enableRequest.Validate()
	assert.Nil(t, err)

	assert.Equal(t, enableRequest.Endpoint(), "/v2/users/123/contacts/1234/disable")
	assert.Equal(t, enableRequest.Method(), "POST")

}