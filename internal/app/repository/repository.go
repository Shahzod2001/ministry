package repository

import (
	"ministry/internal/storage"
	"ministry/pkg/logger"
)

type Repository struct {
	IAuthorization
	ITeacherRepository
	IEntityRepository
}

func New(log *logger.Logger, storage *storage.Storage) *Repository {
	return &Repository{
		IAuthorization:     NewAuthRepository(log, storage),
		ITeacherRepository: NewTeacherRepository(log, storage),
		IEntityRepository:  NewEntityRepository(log, storage),
	}
}
