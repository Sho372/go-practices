package main

import (
	"encoding/json"
	"fmt"
)

type (

	response struct {
		Status int `json:"status"`
		Data data
		Responsetime string
	}

	data struct {
		JpyVolume string
		TierLevel int
		Limit []limit
	}

	limit struct {
		Symbol string
		TodayLimitOpenSize string
		TodayLimitBuySize string
		TodayLimitSellSize string
	}

)

func main() {

	bs := []byte(`{
  "status": 0,
  "data": {
      "jpyVolume": "9988888",
      "tierLevel": 1,
      "limit": [
         {
           "symbol": "BTC/JPY",
           "todayLimitOpenSize": "10000"
         },
         {
           "symbol": "BTC",
           "todayLimitBuySize": "98",
           "todayLimitSellSize": "102"
         }
      ]
  },
  "responsetime": "2019-03-19T02:15:06.055Z"
}`)

//	var res response
	var res interface{}
	err := json.Unmarshal(bs, &res)
	if err != nil {
		panic(nil)
	}
	//%+vでキー名まで表示
	fmt.Printf("%+v\n", res)
}
