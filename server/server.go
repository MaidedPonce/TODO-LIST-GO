package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	db "github.com/MaidedPonce/TODO-LIST-GO/database"
	"github.com/MaidedPonce/TODO-LIST-GO/repository"
	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	DataBaseUrl string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("no hay puerto sjdhsksksjl:D")
	}

	if config.DataBaseUrl == "" {
		return nil, errors.New("no hay url de la db jhassjsjs:(")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := db.NewDBRepository(b.config.DataBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
	log.Println("Ya se inicio la cosa ajjaj:D", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("No hay server:(", err)
	}
}
