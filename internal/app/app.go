package app

import (
	"github.com/22Fariz22/binance-api/internal/config"
	handler "github.com/22Fariz22/binance-api/internal/delivery/http"
	"github.com/22Fariz22/binance-api/internal/usecase"
	"github.com/22Fariz22/binance-api/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type app struct {
	cfg        *config.Config
	httpServer *http.Server
	UC         usecase.UseCase
}

func NewApp(cfg *config.Config) *app {
	return &app{
		cfg:        cfg,
		httpServer: nil,
		UC:         usecase.NewUseCase(),
	}
}

func (a *app) Run() {
	l := logger.New(a.cfg.Log.Level)
	l.Info("app start")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	hd := handler.NewHandler(a.UC, a.cfg, l)

	r.Post("/", hd.GetDiffCurrency)

	http.ListenAndServe(a.cfg.Port, r)
}
