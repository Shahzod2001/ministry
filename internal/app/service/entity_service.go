package service

import "ministry/internal/model"

type IEntityService interface {
	GetAllGenders() ([]model.Sex, error)
	GetAllAcademicDegrees() ([]model.AcademicDegree, error)
	GetAllAcademicPositions() ([]model.AcademicPosition, error)
	GetAllSpecs() ([]model.Spec, error)
	GetAllDirectionSpecs() ([]model.DirectionSpec, error)
	GetAllTeacherTypes() ([]model.TeacherType, error)
	GetAllCities() ([]model.City, error)
}

type EntityService struct {
	rep IEntityService
}

func NewEntityService(rep IEntityService) *EntityService {
	return &EntityService{
		rep: rep,
	}
}

func (e *EntityService) GetAllGenders() ([]model.Sex, error) {
	return e.rep.GetAllGenders()
}

func (e *EntityService) GetAllAcademicDegrees() ([]model.AcademicDegree, error) {
	return e.rep.GetAllAcademicDegrees()
}

func (e *EntityService) GetAllAcademicPositions() ([]model.AcademicPosition, error) {
	return e.rep.GetAllAcademicPositions()
}

func (e *EntityService) GetAllSpecs() ([]model.Spec, error) {
	return e.rep.GetAllSpecs()
}

func (e *EntityService) GetAllDirectionSpecs() ([]model.DirectionSpec, error) {
	return e.rep.GetAllDirectionSpecs()
}

func (e *EntityService) GetAllTeacherTypes() ([]model.TeacherType, error) {
	return e.rep.GetAllTeacherTypes()
}

func (e *EntityService) GetAllCities() ([]model.City, error) {
	return e.rep.GetAllCities()
}
