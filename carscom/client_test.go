package carscom

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {
	client := &CarsComClient{
		Client: &http.Client{},
	}
	listings, err := client.Search(&SearchParams{
		Zip:             "12345",
		MaximumDistance: CustomMaxDistance("50"),
	})
	if err != nil {
		t.Fatalf("Issue searching listings %v\n", err)
	} else if len(listings) <= 0 {
		t.Fatalf("No listings found\n")
	}
	fmt.Printf("Length of listings: %d\n", len(listings))
	fmt.Printf(listings[rand.Intn(len(listings))].Title + "\n")
}

func TestSearchModel(t *testing.T) {
	client := &CarsComClient{
		Client: &http.Client{},
	}
	listings, err := client.Search(&SearchParams{
		Zip:             "12345",
		Models:          []string{"chevrolet-corvette"},
		MaximumDistance: CustomMaxDistance("50"),
	})
	if err != nil {
		t.Fatalf("Issue searching listings %v\n", err)
	} else if len(listings) <= 0 {
		t.Fatalf("No listings found\n")
	}
	fmt.Printf("Length of listings: %d\n", len(listings))
	fmt.Printf(listings[rand.Intn(len(listings))].Title + "\n")
}

func TestSearchPrice(t *testing.T) {
	client := &CarsComClient{
		Client: &http.Client{},
	}
	listings, err := client.Search(&SearchParams{
		Zip:             "12345",
		Models:          []string{"chevrolet-corvette"},
		MaximumDistance: CustomMaxDistance("50"),
		ListMaxPrice:    25000,
	})
	if err != nil {
		t.Fatalf("Issue searching listings %v\n", err)
	} else if len(listings) <= 0 {
		t.Fatalf("No listings found\n")
	}
	fmt.Printf("Length of listings: %d\n", len(listings))
	l := listings[rand.Intn(len(listings))]
	fmt.Println(l.Title, l.PrimaryPrice, l.PriceDrop, l.DealPrice)
}

func TestSearchCondition(t *testing.T) {
	client := &CarsComClient{
		Client: &http.Client{},
	}
	listings, err := client.Search(&SearchParams{
		StockType: ST_ALL,
		Zip:       "12345",
		Models:    []string{"chevrolet-corvette"},
	})
	if err != nil {
		t.Fatalf("Issue searching listings %v\n", err)
	} else if len(listings) <= 0 {
		t.Fatalf("No listings found\n")
	}
	fmt.Printf("Length of listings: %d\n", len(listings))
	l := listings[0]
	fmt.Println(l.Title, l.PrimaryPrice, l.VIN)
}
