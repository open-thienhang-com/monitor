package strategy

import "mono.thienhang.com/pkg/models/base"

type Strategy struct {
	base.Base
	Name        string `json:"name"`
	TimeStart   string `json:"time_start"`
	TimeEnd     string `json:"time_end"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
