package customer

import "mono.thienhang.com/pkg/models/base"

type Customer struct {
	base.Base
	Name          string `json:"name"`
	Gender        string `json:"gender"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	City          string `json:"city"`
	Country       string `json:"country"`
	StreetName    string `json:"street_name"`
	StreetNumber  string `json:"street_number"`
	MarketSegment string `json:"market_segment"`
}
