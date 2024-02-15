package discount

import "mono.thienhang.com/pkg/models/base"

type Discount struct {
	base.Base
	DisplayName string  `json:"display_name"`
	Value       float64 `json:"value"`
	MaxValue    float64 `json:"max_value"`
	IsDelete    bool    `json:"is_delete"`
}
