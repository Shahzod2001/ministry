package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "ministry/docs"
	"ministry/internal/app/handler"
	"ministry/internal/app/handler/middleware"
)

// @title           Swagger Ministry Project
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8585

// @securityDefinitions.basic  BasicAuth
func Setup(h *handler.Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	auth := router.Group("auth")
	{
		auth.POST("sign-up", h.SignUp)
		auth.POST("sign-in", h.SignIn)
		auth.POST("sign-in/admin", h.SignInAdmin)
	}

	teacher := router.Group("teacher", middleware.AuthMW())
	{
		teacher.POST("", h.AddTeacher)
		teacher.PUT(":id", h.EditTeacher)
		teacher.GET("", h.GetUniversityTeachers)
	}

	admin := router.Group("admin", middleware.AuthMW(), middleware.AdminMW())
	{
		admin.GET("teachers", h.GetAllUniversityTeachers)
	}

	entity := router.Group("entity")
	{
		entity.GET("genders", h.GetAllGenders)
		entity.GET("academic-degrees", h.GetAllAcademicDegrees)
		entity.GET("academic-positions", h.GetAllAcademicPositions)
		entity.GET("specs", h.GetAllSpecs)
		entity.GET("direction-specs", h.GetAllDirectionSpecs)
		entity.GET("teacher-types", h.GetAllTeacherTypes)
		entity.GET("cities", h.GetAllCities)
	}

	router.NoRoute(h.NoRoute)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("ping", h.Ping)

	return router
}
