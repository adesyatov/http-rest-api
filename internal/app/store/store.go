package store

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

type Store struct {
	config               *Config
	db                   *sql.DB
	peopleTempRepository *PeopleTempRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("sqlserver", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) PeopleTemp() *PeopleTempRepository {
	if s.peopleTempRepository != nil {
		return s.peopleTempRepository
	}

	s.peopleTempRepository = &PeopleTempRepository{
		store: s,
	}

	return s.peopleTempRepository
}
