package carscom

import "github.com/PuerkitoBio/goquery"

type Listing struct {
	Title        string
	Miles        string
	Link         string
	PrimaryPrice string
	PriceDrop    string
	Carfax       string
	DealerName   string
	MilesFrom    string
	DealLabel    string
	DealPrice    string
	VIN          string
}

func ParseListings(doc *goquery.Document) []Listing {
	var listings []Listing
	doc.Find(".vehicle-card   ").Each(func(i int, selection *goquery.Selection) {
		l := Listing{}

		l.Title = selection.Find(".title").Text()
		l.Miles = selection.Find(".mileage").Text()
		l.Link = selection.Find(".vehicle-card-link").AttrOr("href", "")
		l.PrimaryPrice = selection.Find(".primary-price").Text()
		l.PriceDrop = selection.Find(".secondary-price").Text()
		l.Carfax = selection.Find(".sds-link--ext").AttrOr("href", "")
		l.DealerName = selection.Find(".dealer-name > strong").Text()
		l.MilesFrom = selection.Find(".miles-from ").Text()
		l.DealLabel = selection.Find(".sds-badge__label").Text()
		l.DealPrice = selection.Find(".sds-badge--price-savings").Text()
		l.VIN = selection.Find(".contact-by-phone").AttrOr("data-vim", "")

		listings = append(listings, l)
	})
	return listings
}
