package store

import (
	"github.com/adesyatov/http-rest-api/internal/app/model"
)

type PeopleTempRepository struct {
	store *Store
}

func (r *PeopleTempRepository) Create(p *model.PeopleTemp) (*model.PeopleTemp, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO people_temp (email, encrypted_password) VALUES (@p1, @p2)\nSELECT TOP 1 idpeopletemp FROM people_temp ORDER BY idpeopletemp DESC",
		p.Email,
		p.EncryptedPassword,
	).Scan(&p.Idpeopletemp); err != nil {
		return nil, err
	}

	return p, nil
}

func (r *PeopleTempRepository) FindByEmail(email string) (*model.PeopleTemp, error) {
	p := &model.PeopleTemp{}
	if err := r.store.db.QueryRow(
		"SELECT TOP 1 idpeopletemp, email, encrypted_password FROM people_temp WHERE email = @p1",
		email,
	).Scan(
		&p.Idpeopletemp,
		&p.Email,
		&p.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return p, nil
}
