package brewerydb

import (
	"os"
	"testing"
)

func TestSearchBreweryDB(t *testing.T) {
	query := "bud"
	page := 0
	client := NewClient(os.Getenv("BREWERYDB_KEY"))
	response := client.SearchBeers(query, page)
	if len(response.Beers) == 0 {
		t.Errorf("no beers returned for query '%v'\n", query)
	}
}
