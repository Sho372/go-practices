package main

import (
	//"encoding/json"
	"fmt"
	"log"

	"sho372.tech/http-client-error-handling-practice/apiclient"
	"sho372.tech/http-client-error-handling-practice/models"
)

func main() {

	c := apiclient.New()
	res, err := c.Ticker(models.TickerArgs{Symbol: "BTC"})
	if err != nil {
		log.Fatalf("Error getting ticker data: %v", err)
	}
	fmt.Printf("res: %v", res)

	//	j := []byte(`{"status":2, "messages": [{"message_code": "ERR-5207","message_string":"Not found"}]}`)
	//	var errRes models.ErrorResponse
	//	err2 := json.Unmarshal(j, &errRes)
	//	if err2 != nil {
	//		log.Println(err2.Error())
	//		return
	//	}
	//	fmt.Println(errRes)

}
