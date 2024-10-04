package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetAllGenders(c *gin.Context) {
	var resp response
	genders, err := h.Service.IEntityService.GetAllGenders()
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Success"
	resp.Payload = genders
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllAcademicDegrees(c *gin.Context) {
	var resp response
	acadDegrees, err := h.Service.IEntityService.GetAllAcademicDegrees()
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Success"
	resp.Payload = acadDegrees
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllAcademicPositions(c *gin.Context) {
	var resp response
	acadPositions, err := h.Service.IEntityService.GetAllAcademicPositions()
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Success"
	resp.Payload = acadPositions
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllSpecs(c *gin.Context) {
	var resp response
	specs, err := h.Service.IEntityService.GetAllSpecs()
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Success"
	resp.Payload = specs
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllDirectionSpecs(c *gin.Context) {
	var resp response
	directionSpecs, err := h.Service.IEntityService.GetAllDirectionSpecs()
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Success"
	resp.Payload = directionSpecs
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllTeacherTypes(c *gin.Context) {
	var resp response
	teacherTypes, err := h.Service.IEntityService.GetAllTeacherTypes()
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Success"
	resp.Payload = teacherTypes
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAllCities(c *gin.Context) {
	var resp response
	cities, err := h.Service.IEntityService.GetAllCities()
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Success"
	resp.Payload = cities
	c.JSON(http.StatusOK, resp)
}
