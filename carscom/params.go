package carscom

import (
	"fmt"
	"net/url"
	"strings"
)

type SearchParams struct {
	StockType       StockType
	Makes           []string
	Models          []string
	ListMaxPrice    int
	MaximumDistance MaximumDistance
	Zip             string
	Page            int
	PageSize        int
}

func (sp *SearchParams) Build() string {
	params := url.Values{}
	params.Add("stock_type", string(sp.StockType))
	params.Add("makes[]", strings.Join(sp.Makes, ","))
	params.Add("models[]", strings.Join(sp.Models, ","))
	params.Add("list_price_max", defaultIfZero(sp.ListMaxPrice, ""))
	params.Add("maximum_distance", string(sp.MaximumDistance))
	params.Add("zip", sp.Zip)
	params.Add("page", defaultIfZero(sp.Page, "1"))
	params.Add("page_size", defaultIfZero(sp.PageSize, "20"))
	return params.Encode()
}

func defaultIfZero(i int, def string) string {
	if i == 0 {
		return def
	} else {
		return fmt.Sprintf("%d", i)
	}
}
