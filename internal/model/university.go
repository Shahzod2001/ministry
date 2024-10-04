package model

import "time"

type Gender uint

const (
	Male Gender = iota + 1
	Female
)

type University struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Login     string     `json:"login"`
	Password  string     `json:"password"`
	CityID    int        `json:"city_id"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Teacher struct {
	ID                 int        `json:"id"`
	LastName           string     `json:"last_name"`
	FirstName          string     `json:"first_name" validate:"required"`
	MiddleName         string     `json:"middle_name"`
	BirthDate          time.Time  `json:"birth_date"`
	BirthPlace         string     `json:"birth_place"`
	Gender             Gender     `json:"gender"`
	UniversityID       int        `json:"university_id"`
	AcademicDegreeID   int        `json:"academic_degree_id"`
	AcademicPositionID int        `json:"academic_position_id"`
	SpecID             int        `json:"spec_id"`
	DirectionSpecID    int        `json:"direction_spec_id"`
	TypeID             int        `json:"type_id"`
	JobTitle           string     `json:"job_title"`
	OtherJob           string     `json:"other_job"`
	FromYear           int        `json:"from_year"`
	ToYear             int        `json:"to_year"`
	IsActive           bool       `json:"is_active"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`
}

type AcademicDegree struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AcademicPosition struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Spec struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DirectionSpec struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TeacherType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Sex struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
