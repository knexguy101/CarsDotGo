package autotrader

type listingResponse struct {
	Listings []Listing `json:"listings"`
}

type Listing struct {
	Accelerate          bool     `json:"accelerate"`
	BodyStyleCodes      []string `json:"bodyStyleCodes"`
	CarSchemaOwnerImage string   `json:"carSchemaOwnerImage"`
	DealerConfig        struct {
	} `json:"dealerConfig"`
	Description struct {
		HasMore bool   `json:"hasMore"`
		Label   string `json:"label"`
	} `json:"description"`
	Features        []string    `json:"features"`
	FuelType        string      `json:"fuelType"`
	HasSpecialOffer bool        `json:"hasSpecialOffer"`
	ID              interface{} `json:"id"`
	Images          struct {
		HasVideoIcon bool `json:"hasVideoIcon"`
		Primary      int  `json:"primary"`
		Sources      []struct {
			Alt string `json:"alt"`
			Src string `json:"src"`
		} `json:"sources"`
	} `json:"images"`
	IsHot           bool     `json:"isHot"`
	IsNewlyListed   bool     `json:"isNewlyListed,omitempty"`
	ListingTypes    []string `json:"listingTypes"`
	Make            string   `json:"make"`
	MakeCode        string   `json:"makeCode"`
	Model           string   `json:"model"`
	ModelCode       string   `json:"modelCode"`
	Owner           int      `json:"owner"`
	OwnerName       string   `json:"ownerName"`
	PaymentServices struct {
		CoxAutoDrEnabled bool `json:"coxAutoDrEnabled"`
		DealerSettings   struct {
			AccountStatus            string  `json:"accountStatus"`
			AvailableLeaseTerms      []int   `json:"availableLeaseTerms"`
			AvailableMilesPerYear    []int   `json:"availableMilesPerYear"`
			CashDeal                 bool    `json:"cashDeal"`
			CashDealPrice            int     `json:"cashDealPrice"`
			CreditScoreDefault       string  `json:"creditScoreDefault"`
			DefaultDealType          string  `json:"defaultDealType"`
			DefaultLeaseTerms        int     `json:"defaultLeaseTerms"`
			FinanceCashDownMethod    string  `json:"financeCashDownMethod"`
			FinanceCashDownPct       float64 `json:"financeCashDownPct"`
			FinanceCashDownValue     float64 `json:"financeCashDownValue"`
			LeaseCalculateOnMsrp     bool    `json:"leaseCalculateOnMsrp"`
			LeaseCashDownAmount      int     `json:"leaseCashDownAmount"`
			LeaseCashDownMethod      string  `json:"leaseCashDownMethod"`
			LeaseCashDownPct         float64 `json:"leaseCashDownPct"`
			LeaseDeductDueAtSigning  bool    `json:"leaseDeductDueAtSigning"`
			LeaseEnabled             bool    `json:"leaseEnabled"`
			MmdID                    string  `json:"mmdId"`
			Negotiation              bool    `json:"negotiation"`
			PaymentServicesEnabled   bool    `json:"paymentServicesEnabled"`
			PaymentServicesPartnerID string  `json:"paymentServicesPartnerId"`
			TaxesFeesEnabled         bool    `json:"taxesFeesEnabled"`
			UIEnabled                bool    `json:"uiEnabled"`
		} `json:"dealerSettings"`
		IncentivesURL            string `json:"incentivesUrl"`
		PaymentCalculationURL    string `json:"paymentCalculationUrl"`
		PaymentServiceActive     bool   `json:"paymentServiceActive"`
		PreferredLenderSpecified bool   `json:"preferredLenderSpecified"`
		DigitalRetailingType     string `json:"digitalRetailingType"`
	} `json:"paymentServices"`
	Phone struct {
		LinkText      string `json:"linkText"`
		PrivateNumber bool   `json:"privateNumber"`
		Value         string `json:"value"`
		IsVisible     bool   `json:"isVisible"`
	} `json:"phone"`
	PriceValidUntil string `json:"priceValidUntil"`
	PricingDetail   struct {
		DealerDiscountedPrice int    `json:"dealerDiscountedPrice"`
		Incentive             int    `json:"incentive"`
		Msrp                  int    `json:"msrp"`
		NoPriceLabel          string `json:"noPriceLabel"`
		Payments              struct {
			Label             string `json:"label"`
			PaymentDisclaimer string `json:"paymentDisclaimer"`
			Value             int    `json:"value"`
		} `json:"payments"`
		Reductions []struct {
			Tooltip string `json:"tooltip"`
			Value   string `json:"value"`
			Label   string `json:"label"`
		} `json:"reductions"`
		SalePrice int `json:"salePrice"`
	} `json:"pricingDetail,omitempty"`
	Priority          string   `json:"priority"`
	QuickViewFeatures []string `json:"quickViewFeatures"`
	Specifications    struct {
		InteriorColor struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"interiorColor"`
		Transmission struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"transmission"`
		Color struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"color"`
		Mpg struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"mpg"`
		Engine struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"engine"`
		DriveType struct {
			Label string `json:"label"`
			Value string `json:"value"`
		} `json:"driveType"`
		Mileage struct {
		} `json:"mileage"`
	} `json:"specifications,omitempty"`
	StockNumber string   `json:"stockNumber"`
	Style       []string `json:"style"`
	Title       string   `json:"title"`
	Trim        string   `json:"trim"`
	AtTrim      string   `json:"atTrim"`
	Type        string   `json:"type"`
	Vin         string   `json:"vin"`
	Website     struct {
		Href string `json:"href"`
	} `json:"website"`
	Year     int      `json:"year"`
	Zip      string   `json:"zip"`
	Packages []string `json:"packages,omitempty"`
}
