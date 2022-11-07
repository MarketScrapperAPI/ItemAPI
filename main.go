package main

import "github.com/MarketScrapperAPI/ItemAPI/api"

func main() {

	api := api.NewApi("50005")
	api.Start()
}
