package store

import "mono.thienhang.com/pkg/models/base"

type Store struct {
	base.Base
	Name          string `json:"name"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Street_name   string `json:"street_name"`
	Street_number string `json:"street_number"`
	Manager       string `json:"manager"`
}
