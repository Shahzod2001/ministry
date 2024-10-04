package repository

import (
	"ministry/internal/model"
	"ministry/internal/model/dto"
	"ministry/internal/storage"
	"ministry/pkg/logger"
)

type ITeacherRepository interface {
	CreateTeacher(teacher *model.Teacher) (int, error)
	EditTeacher(teacher *model.Teacher) error
	GetUniversityTeachers(univerId int) ([]dto.Teacher, error)
	GetAllUniversityTeachers() ([]dto.Teacher, error)
}

type TeacherRepository struct {
	log     *logger.Logger
	storage *storage.Storage
}

func NewTeacherRepository(log *logger.Logger, storage *storage.Storage) *TeacherRepository {
	return &TeacherRepository{
		log:     log,
		storage: storage,
	}
}

func (t *TeacherRepository) CreateTeacher(teacher *model.Teacher) (int, error) {
	var id int

	tx, err := t.storage.Postgres.Begin()
	if err != nil {
		return 0, err
	}

	query := `insert into teachers (last_name, first_name, middle_name, birth_date, birth_place, gender, university_id, academic_degree_id, academic_position_id, spec_id, direction_spec_id, type_id, job_title, other_job, from_year, to_year) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) returning id;`

	if err = tx.QueryRow(query, teacher.LastName, teacher.FirstName, teacher.MiddleName, teacher.BirthDate, teacher.BirthPlace, teacher.Gender, teacher.UniversityID, teacher.AcademicDegreeID, teacher.AcademicPositionID, teacher.SpecID, teacher.DirectionSpecID, teacher.TypeID, teacher.JobTitle, teacher.OtherJob, teacher.FromYear, teacher.ToYear).Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (t *TeacherRepository) EditTeacher(teacher *model.Teacher) error {
	tx, err := t.storage.Postgres.Begin()
	if err != nil {
		return err
	}
	query := `update teachers set last_name = $1, first_name = $2, middle_name = $3, birth_date = $4, birth_place = $5, gender = $6, university_id = $7, academic_degree_id = $8, academic_position_id = $9, spec_id = $10, direction_spec_id = $11, type_id = $12, job_title = $13, other_job = $14, from_year = $15, to_year = $16, updated_at = now() where id = $17;`
	_, err = tx.Exec(query, teacher.LastName, teacher.FirstName, teacher.MiddleName, teacher.BirthDate, teacher.BirthPlace, teacher.Gender, teacher.UniversityID, teacher.AcademicDegreeID, teacher.AcademicPositionID, teacher.SpecID, teacher.DirectionSpecID, teacher.TypeID, teacher.JobTitle, teacher.OtherJob, teacher.FromYear, teacher.ToYear, teacher.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (t *TeacherRepository) GetUniversityTeachers(univerId int) ([]dto.Teacher, error) {
	var teachers []dto.Teacher

	query := `select t.id,
       t.last_name,
       t.first_name,
       t.middle_name,
       t.birth_date,
       t.birth_place,
       g.name  as gender,
       u.name  as university,
       ad.name as academic_degree,
       ap.name as academic_position,
       ts.name as spec,
       ds.name as direction_spec,
       tt.name as type,
       t.job_title,
       t.other_job,
       t.from_year,
       t.to_year,
       t.is_active,
       t.created_at,
       t.updated_at,
       t.deleted_at
from teachers t
         left join genders g on g.id = t.gender
         left join universities u on u.id = t.university_id
         left join academic_degrees ad on ad.id = t.academic_degree_id
         left join academic_positions ap on ap.id = t.academic_position_id
         left join teacher_specialities ts on ts.id = t.spec_id
         left join direction_specialities ds on ds.id = t.direction_spec_id
         left join teacher_types tt on tt.id = t.type_id
where u.id = $1 order by t.id;`

	rows, err := t.storage.Postgres.Query(query, univerId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var teacher dto.Teacher
		err = rows.Scan(&teacher.ID, &teacher.LastName, &teacher.FirstName, &teacher.MiddleName, &teacher.BirthDate, &teacher.BirthPlace, &teacher.Gender, &teacher.University, &teacher.AcademicDegree, &teacher.AcademicPosition, &teacher.Spec, &teacher.DirectionSpec, &teacher.Type, &teacher.JobTitle, &teacher.OtherJob, &teacher.FromYear, &teacher.ToYear, &teacher.IsActive, &teacher.CreatedAt, &teacher.UpdatedAt, &teacher.DeletedAt)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

func (t *TeacherRepository) GetAllUniversityTeachers() ([]dto.Teacher, error) {
	var teachers []dto.Teacher

	query := `select t.id,
       t.last_name,
       t.first_name,
       t.middle_name,
       t.birth_date,
       t.birth_place,
       g.name  as gender,
       u.name  as university,
       ad.name as academic_degree,
       ap.name as academic_position,
       ts.name as spec,
       ds.name as direction_spec,
       tt.name as type,
       t.job_title,
       t.other_job,
       t.from_year,
       t.to_year,
       t.is_active,
       t.created_at,
       t.updated_at,
       t.deleted_at
from teachers t
         left join genders g on g.id = t.gender
         left join universities u on u.id = t.university_id
         left join academic_degrees ad on ad.id = t.academic_degree_id
         left join academic_positions ap on ap.id = t.academic_position_id
         left join teacher_specialities ts on ts.id = t.spec_id
         left join direction_specialities ds on ds.id = t.direction_spec_id
         left join teacher_types tt on tt.id = t.type_id order by t.id;`

	rows, err := t.storage.Postgres.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var teacher dto.Teacher
		err = rows.Scan(&teacher.ID, &teacher.LastName, &teacher.FirstName, &teacher.MiddleName, &teacher.BirthDate, &teacher.BirthPlace, &teacher.Gender, &teacher.University, &teacher.AcademicDegree, &teacher.AcademicPosition, &teacher.Spec, &teacher.DirectionSpec, &teacher.Type, &teacher.JobTitle, &teacher.OtherJob, &teacher.FromYear, &teacher.ToYear, &teacher.IsActive, &teacher.CreatedAt, &teacher.UpdatedAt, &teacher.DeletedAt)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}

	return teachers, nil
}
