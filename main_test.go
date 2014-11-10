package brewerydb

import (
	"fmt"
	"os"
	"testing"
)

func TestSearchBreweryDB(t *testing.T) {

	appID := os.Getenv("VINTANK_APP_ID")
	if len(appID) == 0 {
		panic("VINTANK_APP_ID not set, go and register to get one at developer.cruvee.com")
	}
	secret := os.Getenv("VINTANK_SECRET")
	if len(secret) == 0 {
		panic("VINTANK_SECRET not set, go and register to get one at developer.cruvee.com")
	}

	query := "smoking loon"

	c := NewClient(appID, secret)

	resp := c.SearchWines(query, 0)

	if resp.Total == 0 {
		t.Errorf("no wines returned for query '%v' when there should have been\n", query)
	}
	firstWine := resp.Results[0]
	if len(firstWine.Name) == 0 {
		t.Errorf("no name for first result")
	}
	if firstWine.ABV <= 0 {
		t.Errorf("no ABV for first result")
	}
	if len(firstWine.Brand.Name) == 0 {
		t.Errorf("no Brand for first result")
	}
	fmt.Printf("got wine: %v", (firstWine.Name + " (" + firstWine.Brand.Name + ")\n"))
}
