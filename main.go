package brewerydb

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/tideland/goas/v2/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type apiClient struct {
	appID               string
	secret              string
	baseUrl             string
	endpointSearchWines string
}

type SearchResponse struct {
	Page       int `json:"page"`
	ReqPerPage int `json:"rpp"`
	Total      int `json:"total"`
}

type WineSearchResponse struct {
	SearchResponse
	Results []Wine `json:"results"`
}

type Wine struct {
	Name  string  `json:"name"`
	ABV   float32 `json:"ABV"`
	Brand Brand   `json:"brand"`
}

type Brand struct {
	Name string `json:"name"`
}

func NewClient(appID string, secret string) (c *apiClient) {
	c = new(apiClient)
	c.appID = appID
	c.secret = secret
	c.baseUrl = "http://apiv1.cruvee.com"
	c.endpointSearchWines = "/search/wines"
	return c
}

// SearchWines
// NOTE qstring values are time sensitive, this is just an example:
// http://apiv1.cruvee.com/search/wines?q=smoking&appId=<appID>&sig=1f843e4b311689fa3145f76d1e663268&ts=1415632055270095&fmt=json
func (c *apiClient) SearchWines(q string, pg int) (r WineSearchResponse) {

	logger.SetLevel(logger.LevelDebug)

	logger.Debugf("searching with '%v'", q)

	ts := timestamp()
	sig := c.vintankSignatureFor("GET", c.endpointSearchWines)

	// construct url
	v := url.Values{}
	v.Set("appId", c.appID)
	v.Add("ts", ts)   // add premium features even
	v.Add("sig", sig) // if user isn't signed up for them
	v.Add("q", q)
	v.Add("fmt", "json")
	v.Add("page", strconv.Itoa(pg))
	url := c.baseUrl + c.endpointSearchWines + "?" + v.Encode()

	// make request
	data, err := c.makeGetRequest(url)
	if err != nil {
		logger.Errorf("err: %v", err)
		return
	}

	// deserialize
	err = json.Unmarshal(data, &r)
	if err != nil {
		logger.Errorf("unmarshal err: %v\n", err)
		return
	}

	return
}

// http

func (c *apiClient) makeGetRequest(url string) (buf []byte, err error) {
	logger.Infof("performing request: %v\n", url)

	res, err := http.Get(url)
	if err != nil {
		logger.Errorf("err performing request: %v\n", err)
		return
	}
	if res == nil {
		logger.Errorf("response is nil")
		err = errors.New("response is nil")
		return
	}
	buf, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		logger.Errorf("err reading response body: %v\n", err)
	}
	return
}

// library specific util

func (c *apiClient) vintankSignatureFor(method string, endpoint string) (s string) {
	var b bytes.Buffer
	if len(c.appID) == 0 || len(c.secret) == 0 {
		panic("Vintank API appID or secret not set")
	}

	b.WriteString(c.appID + "\n")
	b.WriteString(method + "\n")
	b.WriteString(c.secret + "\n")
	b.WriteString(timestamp() + "\n")
	b.WriteString(endpoint + "\n")
	rawStr := b.String()
	logger.Debugf("constructing sig with: %v\n", rawStr)
	s = md5Hex(rawStr)
	logger.Debugf("constructed sig: %v", s)
	return
}

// general util

func md5Hex(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func timestamp() string {
	nowMS := time.Now().UnixNano() / 1000
	s := strconv.FormatInt(nowMS, 10)
	logger.Debugf("generated timestamp: %v\n", s)
	return s
}
