package cargurus

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type CarGurusClient struct {
	Client *http.Client
}

func (client *CarGurusClient) Search(params *SearchParams) ([]Listing, error) {
	req, _ := http.NewRequest("GET", "https://www.cargurus.com/Cars/inventorylisting/viewDetailsFilterViewInventoryListing.action?"+params.Build(), nil)
	req.Header = map[string][]string{
		"Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"Accept-Language": {"en-US,en;q=0.9"},
		"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"},
		"Host":            {"www.cargurus.com"},
		"Referer":         {"https://www.cargurus.com/Cars/inventorylisting/viewDetailsFilterViewInventoryListing.action?"}, //this fixes the captcha issue
	}
	res, err := client.Client.Do(req)
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
