package repository

import (
	"database/sql"
	"errors"
	"ministry/internal/model"
	"ministry/internal/storage"
	"ministry/pkg/logger"
)

type IEntityRepository interface {
	GetAllGenders() ([]model.Sex, error)
	GetAllAcademicDegrees() ([]model.AcademicDegree, error)
	GetAllAcademicPositions() ([]model.AcademicPosition, error)
	GetAllSpecs() ([]model.Spec, error)
	GetAllDirectionSpecs() ([]model.DirectionSpec, error)
	GetAllTeacherTypes() ([]model.TeacherType, error)
	GetAllCities() ([]model.City, error)
}

type EntityRepository struct {
	log     *logger.Logger
	storage *storage.Storage
}

func NewEntityRepository(log *logger.Logger, storage *storage.Storage) *EntityRepository {
	return &EntityRepository{
		log:     log,
		storage: storage,
	}
}

func (e *EntityRepository) GetAllGenders() ([]model.Sex, error) {
	var genders []model.Sex
	query := `select id, name from genders`
	rows, err := e.storage.Postgres.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("genders not found")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var gender model.Sex
		if err = rows.Scan(&gender.ID, &gender.Name); err != nil {
			return nil, err
		}
		genders = append(genders, gender)
	}
	return genders, nil
}

func (e *EntityRepository) GetAllAcademicDegrees() ([]model.AcademicDegree, error) {
	var acadDegrees []model.AcademicDegree
	query := `select id, name from academic_degrees`
	rows, err := e.storage.Postgres.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("genders not found")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var acadDegree model.AcademicDegree
		if err = rows.Scan(&acadDegree.ID, &acadDegree.Name); err != nil {
			return nil, err
		}
		acadDegrees = append(acadDegrees, acadDegree)
	}
	return acadDegrees, nil
}

func (e *EntityRepository) GetAllAcademicPositions() ([]model.AcademicPosition, error) {
	var acadPositions []model.AcademicPosition
	query := `select id, name from academic_positions`
	rows, err := e.storage.Postgres.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var acadPosition model.AcademicPosition
		if err = rows.Scan(&acadPosition.ID, &acadPosition.Name); err != nil {
			return nil, err
		}
		acadPositions = append(acadPositions, acadPosition)
	}
	return acadPositions, nil
}

func (e *EntityRepository) GetAllSpecs() ([]model.Spec, error) {
	var specs []model.Spec
	query := `select id, name from teacher_specialities`
	rows, err := e.storage.Postgres.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var spec model.Spec
		if err = rows.Scan(&spec.ID, &spec.Name); err != nil {
			return nil, err
		}
		specs = append(specs, spec)
	}
	return specs, nil
}

func (e *EntityRepository) GetAllDirectionSpecs() ([]model.DirectionSpec, error) {
	var directionSpecs []model.DirectionSpec
	query := `select id, name from direction_specialities`
	rows, err := e.storage.Postgres.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var directionSpec model.DirectionSpec
		if err = rows.Scan(&directionSpec.ID, &directionSpec.Name); err != nil {
			return nil, err
		}
		directionSpecs = append(directionSpecs, directionSpec)
	}
	return directionSpecs, nil
}

func (e *EntityRepository) GetAllTeacherTypes() ([]model.TeacherType, error) {
	var teacherTypes []model.TeacherType
	query := `select id, name from teacher_types`
	rows, err := e.storage.Postgres.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var teacherType model.TeacherType
		if err = rows.Scan(&teacherType.ID, &teacherType.Name); err != nil {
			return nil, err
		}
		teacherTypes = append(teacherTypes, teacherType)
	}
	return teacherTypes, nil
}

func (e *EntityRepository) GetAllCities() ([]model.City, error) {
	var cities []model.City
	query := `select id, name from cities`
	rows, err := e.storage.Postgres.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("")
		}
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var city model.City
		if err = rows.Scan(&city.ID, &city.Name); err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}
	return cities, nil
}
