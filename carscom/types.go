package carscom

type StockType string
type MaximumDistance string

const (
	ST_ALL             StockType = "all"
	ST_NEWANDCERTIFIED StockType = "new_cpo"
	ST_NEW             StockType = "new"
	ST_USED            StockType = "used"
	ST_CERTIFIED       StockType = "cpo"

	MD_ALL MaximumDistance = "all"
)

func CustomMaxDistance(s string) MaximumDistance {
	return MaximumDistance(s)
}
