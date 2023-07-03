package binance

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/22Fariz22/binance-api/internal/entity"
	"github.com/22Fariz22/binance-api/pkg"
	"github.com/22Fariz22/binance-api/pkg/logger"
	"io"
	"net/http"
)

func BinanceAPI(ctx context.Context, l logger.Interface, data *entity.Currency) (string, string, error) {
	//получаем сейчас
	//GET /api/v3/ticker/price
	now, err := now(ctx, l, data)
	if err != nil {
		l.Error("no currency data now")
		return "", "", err
	}

	//получаем то что было 24 часа назад
	//GET /api/v3/ticker/24hr
	//example: api/v3/ticker/24hr?symbol=BTCUSDT
	last24hrs, err := last24hrs(ctx, l, data)
	if err != nil {
		l.Error("no currency data last24hrs")
		return "", "", err
	}

	return now, last24hrs, nil
}

func last24hrs(ctx context.Context, l logger.Interface, data *entity.Currency) (string, error) {
	type CurrencyAPI24hrs struct {
		OpenPrice24hrs string `json:"openPrice"`
	}

	//делаем запрос на курс который бьл 24ч назад
	var api24hrs CurrencyAPI24hrs

	request24hrs := fmt.Sprintf("https://api.binance.com/api/v3/ticker/24hr?symbol=%s%s",
		data.TokenFrom, data.TokenTo)
	res, err := http.Get(request24hrs)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return "", pkg.ErrNoCurrencyData
	}

	payload24hrs, err := io.ReadAll(res.Body)
	if err != nil {
		l.Error("can't read api-body request payload24hrs", err)
		return "", pkg.ErrNoCurrencyData
	}

	if err := json.Unmarshal(payload24hrs, &api24hrs); err != nil {
		l.Info("error unmarshall payload24hrs", err)
		return "", pkg.ErrNoCurrencyData
	}

	return api24hrs.OpenPrice24hrs, nil
}

func now(ctx context.Context, l logger.Interface, data *entity.Currency) (string, error) {
	type CurrencyNow struct {
		PriceNow string `json:"price"`
	}

	var currencyNow CurrencyNow

	requestNow := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s%s",
		data.TokenFrom, data.TokenTo)
	res, err := http.Get(requestNow)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return "", pkg.ErrNoCurrencyData
	}

	payloadNow, err := io.ReadAll(res.Body)
	if err != nil {
		l.Error("can't read api-body request payloadNow", err)
		return "", pkg.ErrNoCurrencyData
	}

	if err := json.Unmarshal(payloadNow, &currencyNow); err != nil {
		l.Info("error unmarshall currencyNow", err)
		return "", pkg.ErrNoCurrencyData
	}

	return currencyNow.PriceNow, nil
}
