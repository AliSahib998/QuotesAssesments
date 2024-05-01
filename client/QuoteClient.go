package client

import (
	"fmt"
	"github.com/AliSahib998/QuotesAssesments/model"
	"github.com/go-resty/resty/v2"
)

type IQuoteClient interface {
	GetRandomQuote() (*model.Quote, error)
}

type QuoteClient struct {
	Client *resty.Client
}

//TODO add error handling here
func (q *QuoteClient) GetRandomQuote() (*model.Quote, error) {

	var quote *model.Quote

	_, err := q.Client.R().
		EnableTrace().SetResult(&quote).
		Get(model.QUOTES_BASE_URL)

	if err != nil {
		return nil, err
	}

	return quote, nil
}
