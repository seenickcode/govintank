#A Simple BreweryDB Wrapper for Go

##Features
* Beer Search

##Setup
* optional: set a BREWERYDB_KEY environment variable if you'd like to run the tests

##Example
````
package main

import (
  "fmt"
  bdb "github.com/seenickcode/brewerydb"
)

func main() {
  
  client := bdb.NewClient("YOUR KEY HERE")
  
  // search for a Bud Light beer
  currPage := 0
  for {
    if len(response.Beers) == 0 {
        break
    }

    response := client.SearchBeers("bud light", currPage)
    
    // print results
    for ndx, beer := range response.Beers {
      fmt.Printf("%d: %v\n", ndx, beer.Name)
    }

    currPage++
  } 

  
}
````