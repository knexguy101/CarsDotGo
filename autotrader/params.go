package carfax

import (
	"fmt"
	"net/url"
	"strings"
)

type SearchParams struct {
	ListingTypes  []string
	MakeCodeList  []string
	ModelCodeList []string
	City          string
	State         string
	Zip           string
	NumRecords    int
	SearchRadius  int
	SortBy        string
	Stats         []string
}

func (sp *SearchParams) Build() string {
	params := url.Values{}
	params.Add("listingTypes", defaultIfZeroArr(sp.ListingTypes, "NEW,USED"))
	params.Add("makeCodeList", strings.Join(sp.MakeCodeList, ","))
	params.Add("modelCodeList", strings.Join(sp.ModelCodeList, ","))
	params.Add("city", sp.City)
	params.Add("state", strings.ToUpper(sp.State))
	params.Add("zip", sp.Zip)
	params.Add("location", "[object Object]")
	params.Add("isNewSearch", "true")
	params.Add("marketExtension", "include")
	params.Add("numRecords", defaultIfZero(sp.NumRecords, "25"))
	params.Add("searchRadius", defaultIfZero(sp.SearchRadius, "75"))
	params.Add("showAccelerateBanner", "false")
	params.Add("sortBy", defaultIfEmpty(sp.SortBy, "relevance"))
	params.Add("dma", "[object Object]")
	params.Add("channel", "ATC")
	params.Add("relevanceConfig", "relevance-v2")
	params.Add("stats", defaultIfZeroArr(sp.Stats, "year,derivedprice"))
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
