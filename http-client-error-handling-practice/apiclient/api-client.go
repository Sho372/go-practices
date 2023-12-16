package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"io"

	"sho372.tech/http-client-error-handling-practice/models"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.coin.z.com",
	Path:   "/public/v1/",
}

type Client struct {
	client *http.Client
}

func New() *Client {
	c := &http.Client{Timeout: time.Minute}

	return &Client{
		client: c,
	}
}

/*
Ticker
*/
func (c *Client) Ticker(args models.TickerArgs) (models.TickerRes, error) {
	// set up a reuqest to the ticker endpoint
	endPoint := baseURL.ResolveReference(&url.URL{Path: "ticker"})
	// add Request headers
	req, err := http.NewRequest("GET", endPoint.String(), nil)
	if err != nil {
		return models.TickerRes{}, err
	}
	req.Header.Add("Accept", "application/json")
	// add query params
	req.URL.RawQuery = args.QueryParams().Encode()

	// send the request
	res, err := c.client.Do(req)
	if err != nil {
		return models.TickerRes{}, err
	}

	//defer res.Body.Close()

	var tickerRes models.TickerRes
	bodyBytes, _ := io.ReadAll(io.Reader(res.Body))
	fmt.Printf("bodyBytes: %v", string(bodyBytes))
	json.Unmarshal(bodyBytes, &tickerRes)
	//	if err := json.NewDecoder(res.Body).Decode(&tickerRes); err != nil {
	//		return models.TickerRes{}, err
	//	}

	//closeしたあとに、再度res.Bodyを読み込むのでdeferは使わず、手動でclose
	res.Body.Close()

	status := tickerRes.Status
	fmt.Printf("status: %v\n", status)

	//https://stackoverflow.com/questions/46948050/how-to-read-request-body-twice-in-golang-middleware
	//error responseをデコードするために、再度res.Bodyのバッファを用意
	res.Body = io.NopCloser(io.Reader(bytes.NewBuffer(bodyBytes)))

	switch status {
	case 0:
		return tickerRes, nil //正常終了
	case 1, 2:
		var errRes models.ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return models.TickerRes{}, err
		}
		fmt.Printf("errRes: %v", errRes)
		return models.TickerRes{}, &errRes
	default:
		return models.TickerRes{}, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}
