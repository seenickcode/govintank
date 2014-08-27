# A Solid BreweryDB Wrapper for Go

## Example

````
package main

import (
  "fmt"
  bdb "seenickcode/brewerydb"
)

func main() {
  
  client := bdb.NewClient("YOUR KEY HERE")
  
  // search for a Bud Light beer
  response := client.Search("bud light", 0)

  // print results
  for ndx, item := range response.Data {
    fmt.Printf("%d: %v\n", ndx, item["name"])
  }
}
````