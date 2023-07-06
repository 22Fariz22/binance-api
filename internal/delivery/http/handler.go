package http

import (
	"context"
	"encoding/json"
	"github.com/22Fariz22/binance-api/config"
	"github.com/22Fariz22/binance-api/internal/entity"
	"github.com/22Fariz22/binance-api/internal/usecase"
	"github.com/22Fariz22/binance-api/pkg/logger"
	"io"
	"net/http"
)

// Handler структура хэндлер
type Handler struct {
	UC  usecase.UseCase
	Cfg config.Config
	l   logger.Interface
}

// NewHandler создает хэндлер
func NewHandler(UC usecase.UseCase, cfg *config.Config, l logger.Interface) *Handler {
	return &Handler{
		UC:  UC,
		Cfg: *cfg,
		l:   l,
	}
}

// GetDiffCurrency принимает пару валюты
func (h *Handler) GetDiffCurrency(w http.ResponseWriter, r *http.Request) {
	//вытащить json данные от клиента
	//отправить в usecase в GetAPI
	//дать ответ клиенту данные о ранзнице курса валют

	ctx := context.Background()

	var req entity.Currency

	payload, err := io.ReadAll(r.Body)
	if err != nil {
		h.l.Error("can't read body request", err)
		http.Error(w, "", 500)
	}

	if err := json.Unmarshal(payload, &req); err != nil {
		h.l.Info("error unmarshall", err)
		return
	}

	//отправляем в usecase пару валюты и получаем
	now, diff, err := h.UC.GetAPI(ctx, h.l, &req)
	if err != nil {
		h.l.Error("error in handler:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request Status 400"))
		return
	}

	resp := &entity.DiffCourse{
		Now:         now,
		DifCourse24: diff,
	}

	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(resp)
	if err != nil {
		h.l.Info("error json.Marshal DiffCourse", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Status Internal Server Error 500"))
		return
	}

	w.Write(res)
}
