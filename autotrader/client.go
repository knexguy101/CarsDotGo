package autotrader

import (
	"encoding/json"
	"io"
	"net/http"
)

type AutotraderClient struct {
	Client *http.Client
}

func (client *AutotraderClient) Search(params *SearchParams) ([]Listing, error) {
	req, _ := http.NewRequest("GET", "https://www.autotrader.com/rest/searchresults/base?"+params.Build(), nil)
	req.Header = map[string][]string{
		"User-Agent": {"insomnia/2023.2.2"},
		"Accept":     {"*/*"},
		"Host":       {"helix.carfax.com"},
	}
	res, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var obj listingResponse
	err = json.Unmarshal(b, &obj)
	if err != nil {
		return nil, err
	}

	return obj.Listings, nil
}
