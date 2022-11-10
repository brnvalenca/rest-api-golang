package repository

import (
	"rest-api/golang/exercise/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAddrRepo struct {
	mock.Mock
}

func (addrepo *MockAddrRepo) SaveAddress(addr *entities.Address) error {
	args := addrepo.Called(addr)
	return args.Error(0)
}

func TestSaveAddress(t *testing.T) {

	addrMock := new(MockAddrRepo)

	addr := entities.BuildAddress(1, "1", "x", "b", "50", "r")

	addrMock.On("SaveAddress", addr).Return(nil)

	err := addrMock.SaveAddress(addr)
	if err != nil {
		t.Errorf(err.Error(), "error saving address")
	}

	assert.Nil(t, err)
}
