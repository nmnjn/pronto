package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nmnjn/pronto/config"
)

type App struct {
	Context context.Context
	Router  *gin.Engine
	Config  config.ProntoConfig
}

func New(ctx context.Context, config config.ProntoConfig) *App {

	if config.Production() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	return &App{
		Context: ctx,
		Router:  router,
		Config:  config,
	}
}

func (app *App) StartServer() {
	app.Router.Run(fmt.Sprintf(":%d", app.Config.Port))
}

func (app *App) EnableHealthCheck() {
	app.Router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Healthy!")
	})
}
