package repository

import (
	"database/sql"
	"errors"
	"ministry/internal/model"
	"ministry/internal/storage"
	"ministry/pkg/logger"
)

type IAuthorization interface {
	SignUp(univer *model.University) (int, error)
	SignIn(univer *model.University) (*model.University, error)
	AdminSignIn(admin *model.Admin) (*model.Admin, error)
}

type AuthRepository struct {
	log     *logger.Logger
	storage *storage.Storage
}

func NewAuthRepository(log *logger.Logger, storage *storage.Storage) *AuthRepository {
	return &AuthRepository{
		log:     log,
		storage: storage,
	}
}

func (r *AuthRepository) SignUp(univer *model.University) (int, error) {
	var id int
	tx, err := r.storage.Postgres.Begin()
	if err != nil {
		r.log.Errorf("error starting transaction: %v", err)
		return 0, err
	}

	query := `insert into universities(name, login, password, city_id) values ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(query, univer.Name, univer.Login, univer.Password, univer.CityID).Scan(&id)
	if err != nil {
		r.log.Errorf("error creating university: %v", err)
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *AuthRepository) SignIn(univer *model.University) (*model.University, error) {
	var univerFromDB model.University
	query := `select id, name, login, password, city_id, is_active, created_at, updated_at, deleted_at from universities where updated_at is null and login = $1`
	err := r.storage.Postgres.QueryRow(query, univer.Login).Scan(&univerFromDB.ID, &univerFromDB.Name, &univerFromDB.Login, &univerFromDB.Password, &univerFromDB.CityID, &univerFromDB.IsActive, &univerFromDB.CreatedAt, &univerFromDB.UpdatedAt, &univerFromDB.DeletedAt)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, ErrUniverNotFound
		}
		r.log.Errorf("error getting university from db: %v", err)
		return nil, err
	}
	return &univerFromDB, nil
}

func (r *AuthRepository) AdminSignIn(admin *model.Admin) (*model.Admin, error) {
	var adminFromDB model.Admin
	query := `select id, last_name, first_name, middle_name, login, password from admins where login = $1`
	err := r.storage.Postgres.QueryRow(query, admin.Login).Scan(&adminFromDB.ID, &adminFromDB.LastName, &adminFromDB.FirstName, &adminFromDB.MiddleName, &adminFromDB.Login, &adminFromDB.Password)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, ErrAdminNotFound
		}
		r.log.Errorf("error getting admin from db: %v", err)
		return nil, err
	}
	return &adminFromDB, nil
}
