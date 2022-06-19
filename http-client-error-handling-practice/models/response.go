package models

import (
	"fmt"
	"net/url"
)

type TickerRes struct {
	Status       int
	Data         []ticker
	Responsetime string
}

type ticker struct {
	Ask       string `json:"ask"`
	Bid       string `json:"bid"`
	High      string `json:"high"`
	Last      string `json:"last"`
	Low       string `json:"low"`
	Symbol    string `json:"symbol"`
	Timestamp string `json:"timestamp"`
	Volue     string `json:"volue"`
}

type TickerArgs struct {
	Symbol string
}

// convert TickerArgs to url.Value
func (args TickerArgs) QueryParams() url.Values {

	q := make(url.Values)

	if args.Symbol != "" {
		q.Add("symbol", args.Symbol)
	}

	return q
}

type ErrorResponse struct {
	Status   int            `json:"status"`
	Messages []ErrorMessage `json:"messages"`
}

type ErrorMessage struct {
	MessageCode   string `json:"message_code"`
	MessageString string `json:"message_string"`
}

// implement the error intreface to be treated as error
func (err *ErrorResponse) Error() string {
	//	return fmt.Sprintf("Status: %v, code: %v, message %v", err.Status, err.Messages[0].MessageCode, err.Messages[0].MessageString)
	return fmt.Sprintf("Status: %v", err.Status)
}
