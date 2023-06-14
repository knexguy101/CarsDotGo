package cargurus

import "github.com/PuerkitoBio/goquery"

type Listing struct {
	Title               string
	Price               string
	Miles               string
	DistanceAndLocation string
	DealRating          string
	Dealership          string
	PropertyList        string
}

func ParseListings(doc *goquery.Document) []Listing {
	var listings []Listing
	doc.Find(".tileWrapper").Each(func(i int, selection *goquery.Selection) {
		l := Listing{}

		l.Title = selection.Find(".titleText").Text()
		l.Price = selection.Find(".price > span").Text()
		l.Miles = selection.Find(".mileageText > span:nth-child(2)").Text()
		l.DistanceAndLocation = selection.Find(".distanceAndLocationText > span:nth-child(2)").Text()
		l.DealRating = selection.Find(".dealRating").Text()
		l.Dealership = selection.Find(".priorityDealerNew").Text()
		l.PropertyList = selection.Find(".propertyList").Children().Text()

		listings = append(listings, l)
	})
	return listings
}
