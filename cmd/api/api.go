package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type API interface {
	Run() error
	Shutdown(ctx context.Context) error
}
type api struct {
	*gorm.DB
	srv *http.Server
}

func NewApi(db *gorm.DB, addr string, mode string) API {
	gin.SetMode(mode)
	engine := gin.Default()

	setupRoutes(engine, db)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", addr),
		Handler:      engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &api{DB: db, srv: srv}
}

func (a *api) Run() error {
	return a.srv.ListenAndServe()
}

func (a *api) Shutdown(ctx context.Context) error {
	return a.srv.Shutdown(ctx)
}
