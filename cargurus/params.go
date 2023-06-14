package cargurus

import (
	"fmt"
	"net/url"
	"strings"
)

type SearchParams struct {
	Zip                       string
	InventorySearchWidgetType string
	ShopByTypes               string
	SortDir                   string
	Distance                  int
	SortType                  string
	SelectedEntity            string
}

func (sp *SearchParams) Build() string {
	params := url.Values{}
	params.Add("zip", sp.Zip)
	params.Add("inventorySearchWidgetType", defaultIfEmpty(sp.InventorySearchWidgetType, "AUTO"))
	params.Add("shopByTypes", defaultIfEmpty(sp.ShopByTypes, "NEAR_BY"))
	params.Add("sortDir", defaultIfEmpty(sp.SortDir, "ASC"))
	params.Add("sourceContext", "carGurusHomePageModel")
	params.Add("distance", defaultIfZero(sp.Distance, "50"))
	params.Add("sortType", defaultIfEmpty(sp.SortType, "DEAL_RATING_RPL"))
	params.Add("entitySelectingHelper.selectedEntity", sp.SelectedEntity)
	return params.Encode()
}

func defaultIfEmpty(s string, def string) string {
	if s == "" {
		return def
	} else {
		return s
	}
}

func defaultIfZero(i int, def string) string {
	if i == 0 {
		return def
	} else {
		return fmt.Sprintf("%d", i)
	}
}

func defaultIfZeroArr(i []string, def string) string {
	if len(i) <= 0 {
		return def
	} else {
		return strings.Join(i, ",")
	}
}
