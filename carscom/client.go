package carscom

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type CarsComClient struct {
	Client *http.Client
}

func (client *CarsComClient) Search(params *SearchParams) ([]Listing, error) {
	res, err := client.Client.Get("https://www.cars.com/shopping/results/?" + params.Build())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return ParseListings(doc), nil
}
