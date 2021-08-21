package handler

import (
	"github.com/moneylion-api/bootstrap"
	"github.com/moneylion-api/env"
	"github.com/moneylion-api/repository"
)

// Handler : Singleton Handler
type Handler struct {
	Repository *repository.Repository
}

// New :
func New(bs *bootstrap.Bootstrap) (*Handler, error) {
	hdl := &Handler{}
	hdl.Repository = repository.New(env.Config.Mongo_DB_Name, bs.MgClient)

	return hdl, nil
}
