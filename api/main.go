package api

import (
	"database/sql"
	"net/http"

	v1 "github.com/cyverse-de/workflows/api/v1"
	"github.com/cyverse-de/workflows/common"
	"github.com/cyverse-de/workflows/model"
	"github.com/labstack/echo"
	"gopkg.in/cyverse-de/messaging.v7"
)

// API defines the REST API of the workflows service
type API struct {
	Echo         *echo.Echo
	AMQPSettings *common.AMQPSettings
	AMQPClient   *messaging.Client
	DB           *sql.DB
	Service      string
	Title        string
	Version      string
}

// RootHandler handles GET requests to the / endpoint.
func (a API) RootHandler(ctx echo.Context) error {
	resp := model.RootResponse{
		Service: a.Service,
		Title:   a.Title,
		Version: a.Version,
	}
	return ctx.JSON(http.StatusOK, resp)
}

// RegisterHandlers registers the supported request handlers.
func (a API) RegisterHandlers() {
	a.Echo.GET("/", a.RootHandler)

	// Register the group for API version 1.
	v1Group := a.Echo.Group("/v1")
	v1API := v1.API{
		Echo:         a.Echo,
		Group:        v1Group,
		AMQPSettings: a.AMQPSettings,
		AMQPClient:   a.AMQPClient,
		DB:           a.DB,
		Service:      a.Service,
		Title:        a.Title,
		Version:      a.Version,
	}
	v1API.RegisterHandlers()
}
