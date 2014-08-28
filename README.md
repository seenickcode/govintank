# A Solid BreweryDB Wrapper for Go

## Features

* Beer Search

## TODO

* Tests

## Example

````
package main

import (
  "fmt"
  bdb "github.com/seenickcode/brewerydb"
)

func main() {
  
  client := bdb.NewClient("YOUR KEY HERE")
  
  // search for a Bud Light beer
  response := client.SearchBeers("bud light", 0)

  // print results
  for ndx, beer := range response.Beers {
    fmt.Printf("%d: %v\n", ndx, beer.Name)
  }
}
````