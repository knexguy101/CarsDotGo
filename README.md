# CarsDotGo
An open source reversal of all the car listing sites search APIs. Search Cars.com, Carfax, Autotrader, CarGurus, and more.

## Structure
```go
// All clients take in an http.Client
// You can set this client to whatever you wish, proxy, timeouts, etc.
// YOU MUST INITIALIZE THE HTTP CLIENT WITHIN EVERY CLIENT

import (
  "https://github.com/knexguy101/CarsDotGo/cargurus"
  "net/http"
)

func main() {
  c := &cargurus.CarGurusClient {
    Client: &http.Client{}
  }
}

//READ THE TEST FILES FOR ACTUAL EXAMPLES
```

## Site Parameters
_Most params can be left blank, they are filled in by default values if so_
### Auto Trader
- ListingTypes: NEW,USED
- MakeCodeList: All Caps First 4 Letters of Maker (EX: Chevy is CHEV, Ford is FORD)
- ModelCodeList: All Caps First 4 Letters of Car Type (EX: Corvette is CORV, Camaro is CAMO)
- City: ...
- State: 2 Letter Abrev
- Zip: ...
- NumRecords: # of Results
- SearchRadius: ...
- SortBy: relevance,
- Stats: year,derivedprice (Leave blank if you don't want to mess with it)

### Carfax
- TpQualityThreshold: 150 (not even sure what this does)
- TpPositions: 1,2,3 (again, not sure what this is)
- tpValueBadges: GOOD,GREAT
- Zip: ...
- Radius: ...
- Sort: BEST
- DynamicRadius: true
- Make: Full Maker's name
- Model: Full Model's name
- Certified: false (Certified pre-owned)

### CarGurus
- Zip: ...
- InventorySearchWidgetType: AUTO
- ShopByTypes: NEAR_BY
- SortDir: ASC, DSC, etc.
- Distance: ...
- SortType: DEAL_RATING_RPL
- SelectedEntity: This is a code given to each car type (EX: Corvette is d606), you have to just find this before using it.

### Cars.com
- StockType: all, new_cpo, new, used, cpo
- Makes: lowercase full name of maker
- Models: (make)-(model) all in lowercase (EX: chevrolet-corvette)
- ListMaxPrice: .... (dollar amount, not cents)
- MaximumDistance: either a distance or ALL
- Zip: ...
- Page: ... (defaults to 1)
- PageSize: ... (defaults to 20)
