package model

type Admin struct {
	ID         int    `json:"id"`
	LastName   string `json:"last_name"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type Role uint

const (
	SuperAdminRole Role = iota + 1
	AdminRole
	ModeratorRole
	UserRole
)

func (r Role) Valid() bool {
	switch r {
	case SuperAdminRole, AdminRole, ModeratorRole, UserRole:
		return true
	}
	return false
}
