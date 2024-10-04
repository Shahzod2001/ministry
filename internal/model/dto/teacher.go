package dto

import "time"

type Teacher struct {
	ID               int        `json:"id"`
	LastName         string     `json:"last_name"`
	FirstName        string     `json:"first_name"`
	MiddleName       string     `json:"middle_name"`
	BirthDate        string     `json:"birth_date"`
	BirthPlace       string     `json:"birth_place"`
	Gender           string     `json:"gender"`
	University       string     `json:"university"`
	AcademicDegree   string     `json:"academic_degree"`
	AcademicPosition string     `json:"academic_position"`
	Spec             string     `json:"spec"`
	DirectionSpec    string     `json:"direction_spec"`
	Type             string     `json:"type"`
	JobTitle         string     `json:"job_title"`
	OtherJob         string     `json:"other_job"`
	FromYear         int        `json:"from_year"`
	ToYear           int        `json:"to_year"`
	IsActive         bool       `json:"is_active"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	DeletedAt        *time.Time `json:"deleted_at,omitempty"`
}
