package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/crypto_stats/model"
	"github.com/labstack/echo"

)


// GetCurrency returns the currency
func (ca *CurrencyAPI) GetCurrency(c echo.Context) error {

	symbolID := c.Param("symbol")
	if symbolID == "all" {
		allCurrency := make([]*model.Currency,0)
		for _, s := range ca.SymbolTicks {
			allCurrency = append(allCurrency,s)
		}
		c.JSON(http.StatusOK, allCurrency)
		return nil
	}
	symbol, ok := ca.SymbolStore[symbolID]
	if !ok {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Symbol %s not supported", symbolID))
		return nil
	}
	resp, err := fetchSymbol(symbol)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
	return err
}

func fetchSymbol(symbol model.Symbol) (*model.Currency, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.hitbtc.com/api/2/public/ticker/%s", symbol.ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	obj := &model.Currency{}
	err = json.Unmarshal(body, obj)
	if err != nil {
		return nil, err
	}
	obj.FeeCurrency = symbol.FeeCurrency
	obj.ID = symbol.ID
	obj.Bid = symbol.Bid
	return obj, err
}

