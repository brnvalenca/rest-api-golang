package controllers

import (
	"rest-api/golang/exercise/domain/entities"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/mock"
)

type MockAddressRepo struct {
	mock.Mock
}

type MockKennelRepo struct {
	mock.Mock
}

type MockKennelServ struct {
	mock.Mock
}

// Address Repository Mock

func (addr *MockAddressRepo) Save(address *entities.Address) error {
	args := addr.Called(address)
	return args.Error(0)
}

// Kennel Repository Mock

func (kr *MockKennelRepo) FindAll() ([]entities.Kennel, error) {
	args := kr.Called()
	return args.Get(0).([]entities.Kennel), args.Error(1)
}

func (kr *MockKennelRepo) Save(k *entities.Kennel) (int, error) {
	args := kr.Called(k)
	return args.Int(0), args.Error(1)
}

func (kr *MockKennelRepo) FindById(id string) (*entities.Kennel, error) {
	args := kr.Called(id)
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (kr *MockKennelRepo) Delete(id string) (*entities.Kennel, error) {
	args := kr.Called(id)
	return args.Get(0).(*entities.Kennel), args.Error(1)
}

func (kr *MockKennelRepo) Update(u *entities.Kennel, id string) error {
	args := kr.Called(u, id)
	return args.Error(0)
}

func (kr *MockKennelRepo) CheckIfExists(id string) bool {
	args := kr.Called(id)
	return args.Bool(0)
}

// Kennel Service Mock
