package brewerydb

import(
  "fmt"
  "net/http"
  "net/url"
  "io/ioutil"
  "encoding/json"
  "strconv"
)

const defaultBaseUrl string = "http://api.brewerydb.com/v2"

type breweryDBClient struct {
  apiKey string
  baseUrl string
  VerboseMode bool
}

type BreweryDBResponse struct {
  CurrentPage int
  NumberOfPages int
  TotalResults int `json:"totalResults"`
  Data []map[string]interface{}
}

type Response map[string]string

func NewClient(apiKey string) (c *breweryDBClient) {
  c = new(breweryDBClient)
  c.apiKey = apiKey 
  c.baseUrl = defaultBaseUrl
  c.VerboseMode = false
  return c
}

func (c *breweryDBClient) Search(q string, pg int) (bdbResp BreweryDBResponse) {
  
  // set up query string then url
  v := url.Values{}
  v.Set("key", c.apiKey)
  v.Add("q", q)  
  v.Add("p", strconv.Itoa(pg))
  url := c.baseUrl + "/search?" + v.Encode()
  
  // perform request and convert response to an object
  c.fetchThenUnmarshal(url, &bdbResp)

  // report our search results
  c.log("got %d total results", bdbResp.TotalResults)
  
  return bdbResp
}

func (c *breweryDBClient) fetchThenUnmarshal(url string, bdbResp *BreweryDBResponse) {
  c.log("fetching: %v\n", url)
  res, err := http.Get(url)
  if err != nil {
    fmt.Printf("http err: %v\n", err)
  }
  buf, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Printf("ioutil err: %v\n", err)
  }
  err = json.Unmarshal(buf, &bdbResp)
  if err != nil {
    fmt.Printf("json err: %v\n", err)
  }
}

func (c *breweryDBClient) log(format string, a ...interface{}) {
  if c.VerboseMode {
    fmt.Printf(format, a)
  }
}