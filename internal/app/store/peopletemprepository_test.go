package store_test

import (
	"testing"

	"github.com/adesyatov/http-rest-api/internal/app/model"
	"github.com/adesyatov/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestPeopleTempRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("people_temp")

	p, err := s.PeopleTemp().Create(&model.PeopleTemp{
		Email: "people@example.com",
	})

	assert.NoError(t, err)
	assert.NotNil(t, p)
}

func TestPeopleTempRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("people_temp")

	email := "people@example.com"
	_, err := s.PeopleTemp().FindByEmail(email)
	assert.Error(t, err)

	s.PeopleTemp().Create(&model.PeopleTemp{
		Email: email,
	})
	p, err := s.PeopleTemp().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, p)
}
