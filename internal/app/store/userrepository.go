package store

import (
	"github.com/adesyatov/http-rest-api/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO [user] (email, encrypted_password) VALUES (@p1, @p2)\nSELECT TOP 1 iduser FROM [user] ORDER BY iduser DESC",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.Iduser); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT TOP 1 iduser, email, encrypted_password FROM [user] WHERE email = @p1",
		email,
	).Scan(
		&u.Iduser,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}

	return u, nil
}
