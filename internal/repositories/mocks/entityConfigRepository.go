package mocks

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"

	"github.com/Permify/permify/internal/repositories/entities"
	"github.com/Permify/permify/pkg/errors"
)

// EntityConfigRepository is an autogenerated mock type for the EntityConfigRepository type
type EntityConfigRepository struct {
	mock.Mock
}

func (_m *EntityConfigRepository) Migrate() errors.Error {
	return nil
}

// All -
func (_m *EntityConfigRepository) All(ctx context.Context, version string) (configs entities.EntityConfigs, err errors.Error) {
	ret := _m.Called(version)

	var r0 []entities.EntityConfig
	if rf, ok := ret.Get(0).(func(context.Context, string) []entities.EntityConfig); ok {
		r0 = rf(ctx, version)
	} else {
		r0 = ret.Get(0).([]entities.EntityConfig)
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.Error); ok {
		r1 = rf(ctx, version)
	} else {
		r1 = ret.Get(1).(errors.Error)
	}

	return r0, r1
}

// Replace -
func (_m *EntityConfigRepository) Read(ctx context.Context, entityName string, version string) (entities.EntityConfig, errors.Error) {
	ret := _m.Called(entityName, version)

	var r0 entities.EntityConfig
	if rf, ok := ret.Get(0).(func(context.Context, string, string) entities.EntityConfig); ok {
		r0 = rf(ctx, entityName, version)
	} else {
		r0 = ret.Get(0).(entities.EntityConfig)
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) errors.Error); ok {
		r1 = rf(ctx, entityName, version)
	} else {
		r1 = ret.Get(1).(errors.Error)
	}

	return r0, r1
}

// Write -
func (_m *EntityConfigRepository) Write(ctx context.Context, configs entities.EntityConfigs, version string) (err errors.Error) {
	ret := _m.Called(configs, version)

	var r0 errors.Error
	if rf, ok := ret.Get(1).(func(context.Context, entities.EntityConfigs, string) errors.Error); ok {
		r0 = rf(ctx, configs, version)
	} else {
		r0 = ret.Get(1).(errors.Error)
	}

	return r0
}
