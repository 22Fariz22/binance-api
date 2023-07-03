package usecase

import (
	"context"
	"fmt"
	"github.com/22Fariz22/binance-api/internal/entity"
)

type UseCase interface {
	GetAPI(ctx context.Context, data *entity.Currency)
}

type useCase struct {
}

// NewUseCase create usecase
func NewUseCase() *useCase {
	return &useCase{}
}

func (u *useCase) GetAPI(ctx context.Context, data *entity.Currency) {
	//создать json и сделать запрос в binance API
	//вернуть ответ клиенту
	fmt.Println(data)

}
