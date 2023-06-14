package cargurus

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {
	client := &CarGurusClient{
		Client: &http.Client{},
	}
	listings, err := client.Search(&SearchParams{
		Zip:            "12345",
		SelectedEntity: "d606",
	})
	if err != nil {
		t.Fatalf("Issue searching listings %v\n", err)
	} else if len(listings) <= 0 {
		t.Fatalf("No listings found\n")
	}
	fmt.Printf("Length of listings: %d\n", len(listings))
	fmt.Printf(listings[rand.Intn(len(listings))].Title + "\n")
}
