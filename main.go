package brewerydb

import (
	// "encoding/json"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/tideland/goas/v2/logger"

	"strconv"

	"time"
	// "io/ioutil"
	// "net/http"
	// "net/url"
	// "strconv"
)

type apiClient struct {
	appID   string
	secret  string
	baseUrl string
}

type SearchResponse struct {
	TotalResults int
}

func NewClient(appID string, secret string) (c *apiClient) {
	c = new(apiClient)
	c.appID = appID
	c.secret = secret
	c.baseUrl = "http://apiv1.cruvee.com"
	return c
}

func (c *apiClient) Search(q string) (r SearchResponse) {

	logger.SetLevel(logger.LevelDebug)

	logger.Debugf("searching with '%v'", q)
	logger.Debugf("generated sig: %v", c.signatureFor("GET", "/search/wines"))

	r.TotalResults = 1

	return
}

func (c *apiClient) signatureFor(method string, endpoint string) (s string) {
	var b bytes.Buffer
	if len(c.appID) == 0 || len(c.secret) == 0 {
		panic("Vintank API appID or secret not set")
	}
	nowMS := time.Now().UnixNano() / 1000

	b.WriteString(c.appID + "\n")
	b.WriteString(method + "\n")
	b.WriteString(c.secret + "\n")
	b.WriteString(strconv.FormatInt(nowMS, 10) + "\n")
	b.WriteString(endpoint + "\n")
	rawStr := b.String()
	logger.Debugf("raw sig: %v\n", rawStr)
	s = md5Hex(rawStr)
	return
}

func md5Hex(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

// func (c *breweryDBClient) performGetRequest(url string) (buf []byte, err error) {
// 	c.log("fetching: %v\n", url)
// 	res, err := http.Get(url)
// 	if err != nil {
// 		fmt.Printf("http err: %v\n", err)
// 		return
// 	}
// 	if res == nil {
// 		err = fmt.Errorf("err, response is nil")
// 		return
// 	}
// 	buf, err = ioutil.ReadAll(res.Body)
// 	defer res.Body.Close()
// 	if err != nil {
// 		fmt.Printf("ioutil err: %v\n", err)
// 	}
// 	return
// }
