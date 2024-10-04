package handler

import (
	"github.com/gin-gonic/gin"
	"ministry/internal/model"
	"net/http"
)

func (h *Handler) SignUp(c *gin.Context) {
	var resp response
	type signUpInput struct {
		Name     string `json:"name" binding:"required"`
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
		CityId   int    `json:"city_id" binding:"required"`
	}

	var input signUpInput
	err := c.BindJSON(&input)
	if err != nil {
		respondWithError(c, &resp, http.StatusBadRequest, err)
		return
	}

	univer := model.University{
		Name:     input.Name,
		Login:    input.Login,
		Password: input.Password,
		CityID:   input.CityId,
	}

	uniVerID, err := h.Service.IAuthorization.SignUp(&univer)
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}
	resp.Code = http.StatusCreated
	resp.Message = "University is created. Please wait for the activation"
	resp.Payload = uniVerID
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) SignIn(c *gin.Context) {
	var resp response
	type signInInput struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var input signInInput
	err := c.BindJSON(&input)
	if err != nil {
		respondWithError(c, &resp, http.StatusBadRequest, err)
		return
	}
	univer := model.University{
		Login:    input.Login,
		Password: input.Password,
	}

	tokenPair, err := h.Service.IAuthorization.SignIn(&univer)
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "Success login"
	resp.Payload = tokenPair
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) SignInAdmin(c *gin.Context) {
	var resp response
	type signInInput struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var input signInInput
	err := c.BindJSON(&input)
	if err != nil {
		respondWithError(c, &resp, http.StatusBadRequest, err)
		return
	}
	admin := model.Admin{
		Login:    input.Login,
		Password: input.Password,
	}

	tokenPair, err := h.Service.IAuthorization.AdminSignIn(&admin)
	if err != nil {
		respondWithError(c, &resp, http.StatusInternalServerError, err)
		return
	}

	resp.Code = http.StatusOK
	resp.Message = "Success login"
	resp.Payload = tokenPair
	c.JSON(http.StatusOK, resp)
}
