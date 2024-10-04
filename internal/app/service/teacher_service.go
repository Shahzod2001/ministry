package service

import (
	"ministry/internal/app/repository"
	"ministry/internal/model"
	"ministry/internal/model/dto"
)

type ITeacherService interface {
	CreateTeacher(teacher *model.Teacher) (int, error)
	EditTeacher(teacher *model.Teacher) error
	GetUniversityTeachers(univerId int) ([]dto.Teacher, error)
	GetAllUniversityTeachers() ([]dto.Teacher, error)
}

type TeacherService struct {
	rep repository.ITeacherRepository
}

func NewTeacherService(rep repository.ITeacherRepository) *TeacherService {
	return &TeacherService{rep: rep}
}

func (t *TeacherService) CreateTeacher(teacher *model.Teacher) (int, error) {
	return t.rep.CreateTeacher(teacher)
}

func (t *TeacherService) EditTeacher(teacher *model.Teacher) error {
	return t.rep.EditTeacher(teacher)
}

func (t *TeacherService) GetUniversityTeachers(univerId int) ([]dto.Teacher, error) {
	return t.rep.GetUniversityTeachers(univerId)
}

func (t *TeacherService) GetAllUniversityTeachers() ([]dto.Teacher, error) {
	return t.rep.GetAllUniversityTeachers()
}
