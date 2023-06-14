package carfax

type listingResponse struct {
	Listings []Listing `json:"listings"`
}

type Listing struct {
	Dealer struct {
		CarfaxID                string  `json:"carfaxId"`
		DealerInventoryURL      string  `json:"dealerInventoryUrl"`
		CfxMicrositeURL         string  `json:"cfxMicrositeUrl"`
		Name                    string  `json:"name"`
		Address                 string  `json:"address"`
		City                    string  `json:"city"`
		State                   string  `json:"state"`
		Zip                     string  `json:"zip"`
		Phone                   string  `json:"phone"`
		Latitude                string  `json:"latitude"`
		Longitude               string  `json:"longitude"`
		DealerAverageRating     float64 `json:"dealerAverageRating"`
		DealerReviewComments    string  `json:"dealerReviewComments"`
		DealerReviewDate        string  `json:"dealerReviewDate"`
		DealerReviewReviewer    string  `json:"dealerReviewReviewer"`
		DealerReviewRating      int     `json:"dealerReviewRating"`
		DealerReviewCount       int     `json:"dealerReviewCount"`
		DdcValue                float64 `json:"ddcValue"`
		DealerBadgingExperience string  `json:"dealerBadgingExperience"`
	} `json:"dealer,omitempty"`
	ID                     string   `json:"id"`
	Vin                    string   `json:"vin"`
	Year                   int      `json:"year"`
	Make                   string   `json:"make"`
	Model                  string   `json:"model"`
	Trim                   string   `json:"trim"`
	SubTrim                string   `json:"subTrim"`
	TopOptions             []string `json:"topOptions"`
	Mileage                int      `json:"mileage"`
	ListPrice              int      `json:"listPrice"`
	CurrentPrice           int      `json:"currentPrice"`
	MonthlyPaymentEstimate struct {
		Price              int     `json:"price"`
		DownPaymentPercent int     `json:"downPaymentPercent"`
		InterestRate       int     `json:"interestRate"`
		TermInMonths       int     `json:"termInMonths"`
		LoanAmount         float64 `json:"loanAmount"`
		DownPaymentAmount  float64 `json:"downPaymentAmount"`
		MonthlyPayment     float64 `json:"monthlyPayment"`
	} `json:"monthlyPaymentEstimate"`
	Badge            string `json:"badge,omitempty"`
	ExteriorColor    string `json:"exteriorColor"`
	InteriorColor    string `json:"interiorColor"`
	Engine           string `json:"engine"`
	Displacement     string `json:"displacement"`
	Drivetype        string `json:"drivetype"`
	Transmission     string `json:"transmission"`
	Fuel             string `json:"fuel"`
	MpgCity          int    `json:"mpgCity"`
	MpgHighway       int    `json:"mpgHighway"`
	Bodytype         string `json:"bodytype"`
	VehicleCondition string `json:"vehicleCondition"`
	CabType          string `json:"cabType"`
	BedLength        string `json:"bedLength"`
	FollowCount      int    `json:"followCount"`
	StockNumber      string `json:"stockNumber"`
	ImageCount       int    `json:"imageCount"`
	Images           struct {
		BaseURL    string   `json:"baseUrl"`
		Large      []string `json:"large"`
		Medium     []string `json:"medium"`
		Small      []string `json:"small"`
		FirstPhoto struct {
			Large  string `json:"large"`
			Medium string `json:"medium"`
			Small  string `json:"small"`
		} `json:"firstPhoto"`
	} `json:"images"`
	FirstSeen    string `json:"firstSeen"`
	OneOwner     bool   `json:"oneOwner,omitempty"`
	OwnerHistory struct {
		Text    string `json:"text"`
		Icon    string `json:"icon"`
		IconURL string `json:"iconUrl"`
		History []struct {
			OwnerNumber      int    `json:"ownerNumber"`
			PurchaseDate     string `json:"purchaseDate"`
			EndOwnershipDate string `json:"endOwnershipDate"`
			City             string `json:"city"`
			State            string `json:"state"`
		} `json:"history"`
		IconName string `json:"iconName"`
	} `json:"ownerHistory"`
	NoAccidents     bool `json:"noAccidents,omitempty"`
	AccidentHistory struct {
		Text            string   `json:"text"`
		Icon            string   `json:"icon"`
		IconURL         string   `json:"iconUrl"`
		AccidentSummary []string `json:"accidentSummary"`
		IconName        string   `json:"iconName"`
	} `json:"accidentHistory"`
	ServiceRecords bool `json:"serviceRecords,omitempty"`
	ServiceHistory struct {
		Text    string `json:"text"`
		Icon    string `json:"icon"`
		IconURL string `json:"iconUrl"`
		Number  int    `json:"number"`
		History []struct {
			City            string `json:"city"`
			State           string `json:"state"`
			OdometerReading int    `json:"odometerReading"`
			Date            string `json:"date"`
			Description     string `json:"description"`
			Source          string `json:"source"`
		} `json:"history"`
		IconName string `json:"iconName"`
	} `json:"serviceHistory,omitempty"`
	PersonalUse       bool `json:"personalUse,omitempty"`
	VehicleUseHistory struct {
		Text    string `json:"text"`
		Icon    string `json:"icon"`
		IconURL string `json:"iconUrl"`
		History []struct {
			AverageMilesPerYear int    `json:"averageMilesPerYear"`
			UseType             string `json:"useType"`
			OwnerNumber         int    `json:"ownerNumber"`
		} `json:"history"`
		IconName string `json:"iconName"`
	} `json:"vehicleUseHistory,omitempty"`
	DistanceToDealer        float64  `json:"distanceToDealer"`
	RecordType              string   `json:"recordType"`
	DealerType              string   `json:"dealerType"`
	Advantage               bool     `json:"advantage"`
	VdpURL                  string   `json:"vdpUrl"`
	WindowSticker           string   `json:"windowSticker,omitempty"`
	SortScore               float64  `json:"sortScore"`
	BaseScore               float64  `json:"baseScore"`
	TpCostPerVdp            float64  `json:"tpCostPerVdp,omitempty"`
	AtomOtherOptions        []string `json:"atomOtherOptions"`
	AtomTopOptions          []string `json:"atomTopOptions"`
	TpRetentionScore        float64  `json:"tpRetentionScore,omitempty"`
	DealerBadgingExperience string   `json:"dealerBadgingExperience"`
}
