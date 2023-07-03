package usecase

import (
	"context"
	"github.com/22Fariz22/binance-api/internal/entity"
	"github.com/22Fariz22/binance-api/pkg/binance"
	"github.com/22Fariz22/binance-api/pkg/logger"
	"strconv"
)

type UseCase interface {
	GetAPI(ctx context.Context, l logger.Interface, data *entity.Currency) (float64, float64, error)
}

type useCase struct {
}

// NewUseCase create usecase
func NewUseCase() *useCase {
	return &useCase{}
}

func (u *useCase) GetAPI(ctx context.Context, l logger.Interface, data *entity.Currency) (float64, float64, error) {
	now, last24hrs, err := binance.BinanceAPI(ctx, l, data)
	if err != nil {
		return 0, 0, err
	}

	nowFloat, err := strconv.ParseFloat(now, 32)
	if err != nil {
		return 0, 0, err
	}

	last24hrsFloat, _ := strconv.ParseFloat(last24hrs, 32)

	diff := nowFloat - last24hrsFloat

	return nowFloat, diff, nil
}
