package discount

import "mono.thienhang.com/pkg/models/base"

type DiscountGroup struct {
	base.Base
	QualityDiscount       int     `json:"quality_discount"`
	InvoiceValueCondition float64 `json:"invoice_value_condition"`
	QuantityCondition     int     `json:"quantity_condition"`
	Brand                 string  `json:"brand"`
	ProductID             int     `json:"product_id"`
	Description           string  `json:"description"`
	IsDelete              bool    `json:"is_delete"`
}
