package brewerydb

import (
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

	query := "guinness"

	c := NewClient(appID, secret)

	resp := c.Search(query)

	if resp.TotalResults == 0 {
		t.Errorf("no wines returned for query '%v' when there should have been\n", query)
	}
}
