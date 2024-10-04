package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"ministry/config"
	"ministry/internal/app/handler"
	"ministry/internal/app/repository"
	"ministry/internal/app/router"
	"ministry/internal/app/service"
	"ministry/internal/storage"
	"ministry/pkg/logger"
	"net"
)

type App struct {
	r *repository.Repository
	s *service.Service
	h *handler.Handler
	g *gin.Engine
}

func New(cfg *config.Config, log *logger.Logger) *App {
	app := new(App)
	store := storage.New(log, cfg)
	app.r = repository.New(log, store)
	app.s = service.New(log, app.r)
	app.h = handler.New(app.s)
	app.g = router.Setup(app.h)
	return app
}

func (a *App) Run(cfg *config.Config) {
	address := cfg.Server
	log.Fatal(a.g.Run(net.JoinHostPort(address.Host, address.Port)))
}
