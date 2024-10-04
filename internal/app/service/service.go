package service

import (
	"ministry/internal/app/repository"
	"ministry/pkg/logger"
)

type Service struct {
	IAuthorization
	ITeacherService
	IEntityService
}

func New(log *logger.Logger, rep *repository.Repository) *Service {
	return &Service{
		IAuthorization:  NewAuthService(rep),
		ITeacherService: NewTeacherService(rep),
		IEntityService:  NewEntityService(rep),
	}
}
