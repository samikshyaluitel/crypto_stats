package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/crypto_stats/model"
)


type CurrencyAPI struct{
	Client *http.Client
	SymbolStore map[string]model.Symbol
	SymbolTicks map[string]*model.Currency
}
type Param struct {
	Symbol string `json:"symbol"`
}


func NewCurrencyAPI(cl *http.Client,supportedSymbols []string) (*CurrencyAPI, error){
	if cl == nil{
		cl = http.DefaultClient
	}
	
    keys ,err := fetchSymbols(cl)
	if err != nil {
		return nil, err
	}
	symbolCheck := make(map[string]bool)
	for _,item := range supportedSymbols{
		symbolCheck[item] = true
	} 
	symbolStore := make(map[string]model.Symbol)
	symbolTicks := make(map[string]*model.Currency)

	for _,item := range keys {
		if (symbolCheck[item.ID]){
			symbolStore[item.ID] = item
			var currency = &model.Currency{ID: item.ID,FeeCurrency: item.FeeCurrency}
			go subscribeTicker(item,currency)
			symbolTicks[item.ID] = currency
		}
	}
	
	cAPI := CurrencyAPI{Client: cl,SymbolStore: symbolStore,SymbolTicks: symbolTicks}
	return &cAPI,nil
}

func fetchSymbols(cl *http.Client) ([]model.Symbol,error){
	resp,err := cl.Get("https://api.hitbtc.com/api/2/public/symbol")
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
    keys := make([]model.Symbol,0)
	err = json.Unmarshal(body, &keys)
	
	return keys,err
}


func subscribeTicker(symbol model.Symbol,cur *model.Currency){
	// connect to server
	cl,err := NewWSClient()
	if err != nil {
		panic(err)
	}
	defer cl.Close()
	resCH,err := cl.SubscribeTicker(symbol.ID)
	if err != nil{
		panic(err)
	}
	for i := range resCH {
		// Update the currency on each notification
		cur.Bid=i.Bid
		cur.Ask = i.Ask                  
		cur.Last = i.Last        
		cur.Open=i.Open        
		cur.Low=i.Low         
		cur.High=i.High
	}
}