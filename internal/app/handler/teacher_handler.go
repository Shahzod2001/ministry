package handler

import (
	"github.com/gin-gonic/gin"
	"ministry/internal/model"
	"net/http"
	"strconv"
	"time"
)

// @Summary Add a new teacher
// @Description Add a new teacher to the university
// @Tags teachers
// @Accept json
// @Produce json
// @Param teacher body addTeacherInput true "Add Teacher"
// @Success 200 {object} response
// @Failure 400 {object} response
// @Failure 500 {object} response
// @Router /teacher/create [post]
func (h *Handler) AddTeacher(c *gin.Context) {
	var resp response

	type addTeacherInput struct {
		LastName           string `json:"last_name" binding:"required"`
		FirstName          string `json:"first_name" binding:"required"`
		MiddleName         string `json:"middle_name"`
		BirthDate          string `json:"birth_date" binding:"required"`
		BirthPlace         string `json:"birth_place" binding:"required"`
		Gender             int    `json:"gender" binding:"required"`
		AcademicDegreeID   int    `json:"academic_degree_id" binding:"required"`
		AcademicPositionID int    `json:"academic_position_id" binding:"required"`
		SpecID             int    `json:"spec_id" binding:"required"`
		DirectionSpecID    int    `json:"direction_spec_id" binding:"required"`
		TypeID             int    `json:"type_id" binding:"required"`
		JobTitle           string `json:"job_title" binding:"required"`
		OtherJob           string `json:"other_job"`
		FromYear           int    `json:"from_year" binding:"required"`
		ToYear             int    `json:"to_year" binding:"required"`
	}

	var input addTeacherInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		respondWithError(c, &resp, http.StatusBadRequest, err)
		return
	}

	birthDate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		respondWithError(c, &resp, http.StatusBadRequest, err)
		return
	}

	univerID := c.GetInt("univer_id")

	teacher := model.Teacher{
		LastName:           input.LastName,
		FirstName:          input.FirstName,
		MiddleName:         input.MiddleName,
		BirthDate:          birthDate,
		BirthPlace:         input.BirthPlace,
		Gender:             model.Gender(input.Gender),
		UniversityID:       univerID,
		AcademicDegreeID:   input.AcademicDegreeID,
		AcademicPositionID: input.AcademicPositionID,
		SpecID:             input.SpecID,
		DirectionSpecID:    input.DirectionSpecID,
		TypeID:             input.TypeID,
		JobTitle:           input.JobTitle,
		OtherJob:           input.OtherJob,
		FromYear:           input.FromYear,
		ToYear:             input.ToYear,
	}

	teacherId, err := h.Service.ITeacherService.CreateTeacher(&teacher)
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Teacher successfully added"
	resp.Payload = teacherId
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) EditTeacher(c *gin.Context) {
	var resp response

	type editTeacherInput struct {
		LastName           string `json:"last_name" binding:"required"`
		FirstName          string `json:"first_name" binding:"required"`
		MiddleName         string `json:"middle_name"`
		BirthDate          string `json:"birth_date" binding:"required"`
		BirthPlace         string `json:"birth_place" binding:"required"`
		Gender             int    `json:"gender" binding:"required"`
		AcademicDegreeID   int    `json:"academic_degree_id" binding:"required"`
		AcademicPositionID int    `json:"academic_position_id" binding:"required"`
		SpecID             int    `json:"spec_id" binding:"required"`
		DirectionSpecID    int    `json:"direction_spec_id" binding:"required"`
		TypeID             int    `json:"type_id" binding:"required"`
		JobTitle           string `json:"job_title" binding:"required"`
		OtherJob           string `json:"other_job"`
		FromYear           int    `json:"from_year" binding:"required"`
		ToYear             int    `json:"to_year" binding:"required"`
	}

	var input editTeacherInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		respondWithError(c, &resp, http.StatusBadRequest, err)
		return
	}

	birthDate, err := time.Parse("2006-01-02", input.BirthDate)
	if err != nil {
		respondWithError(c, &resp, http.StatusBadRequest, err)
		return
	}

	teacherId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		respondWithError(c, &resp, http.StatusBadRequest, err)
		return
	}

	univerID := c.GetInt("univer_id")

	teacher := model.Teacher{
		ID:                 teacherId,
		LastName:           input.LastName,
		FirstName:          input.FirstName,
		MiddleName:         input.MiddleName,
		BirthDate:          birthDate,
		BirthPlace:         input.BirthPlace,
		Gender:             model.Gender(input.Gender),
		UniversityID:       univerID,
		AcademicDegreeID:   input.AcademicDegreeID,
		AcademicPositionID: input.AcademicPositionID,
		SpecID:             input.SpecID,
		DirectionSpecID:    input.DirectionSpecID,
		TypeID:             input.TypeID,
		JobTitle:           input.JobTitle,
		OtherJob:           input.OtherJob,
		FromYear:           input.FromYear,
		ToYear:             input.ToYear,
	}

	if err = h.Service.ITeacherService.EditTeacher(&teacher); err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Teacher successfully edited"
	resp.Payload = teacher
	c.JSON(http.StatusOK, resp)
}

// @Summary Get teachers of the university
// @Description Get all teachers associated with the university
// @Tags teachers
// @Accept json
// @Produce json
// @Success 200 {object} response
// @Failure 500 {object} response
// @Router /teacher/all [get]
func (h *Handler) GetUniversityTeachers(c *gin.Context) {
	var resp response

	univerID := c.GetInt("univer_id")

	teachers, err := h.Service.ITeacherService.GetUniversityTeachers(univerID)
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Teachers"
	resp.Payload = teachers

	c.JSON(http.StatusOK, resp)
}

// @Summary Get all teachers
// @Description Get all teachers from all universities
// @Tags teachers
// @Accept json
// @Produce json
// @Success 200 {object} response
// @Failure 500 {object} response
// @Router /admin/teachers [get]
func (h *Handler) GetAllUniversityTeachers(c *gin.Context) {
	var resp response

	teachers, err := h.Service.ITeacherService.GetAllUniversityTeachers()
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Teachers"
	resp.Payload = teachers

	c.JSON(http.StatusOK, resp)
}
