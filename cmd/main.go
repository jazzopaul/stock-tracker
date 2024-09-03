package main

import (
	"github.com/jazzopaul/stock-tracker/alpha_vantage"
)

func main() {

	alphaVantageSvc := alpha_vantage.NewService()
	alphaVantageSvc.ApiCall()
}
