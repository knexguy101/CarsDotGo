package carfax

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type SearchParams struct {
	TpQualityThreshold int
	TpPositions        []string
	TpValueBadges      []string
	Zip                string
	Radius             int
	Sort               string
	DynamicRadius      bool
	Make               string
	Model              string
	Certified          bool
}

func (sp *SearchParams) Build() string {
	params := url.Values{}
	params.Add("tpQualityThreshold", defaultIfZero(sp.TpQualityThreshold, "150"))
	params.Add("tpPositions", defaultIfZeroArr(sp.TpPositions, "1,2,3"))
	params.Add("tpValueBadges", defaultIfZeroArr(sp.TpValueBadges, "GOOD,GREAT"))
	params.Add("zip", sp.Zip)
	params.Add("radius", defaultIfZero(sp.Radius, "50"))
	params.Add("sort", defaultIfEmpty(sp.Sort, "BEST"))
	params.Add("dynamicRadius", strconv.FormatBool(sp.DynamicRadius))
	params.Add("make", sp.Make)
	params.Add("model", sp.Model)
	params.Add("certified", strconv.FormatBool(sp.Certified))
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
