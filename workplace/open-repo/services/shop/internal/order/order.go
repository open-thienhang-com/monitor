package order

import "mono.thienhang.com/pkg/models/base"

type Order struct {
	base.Base
	CustomerID       int     `json:"customer_id"`
	ProductID        int     `json:"product_id"`
	StoreID          int     `json:"store_id"`
	EmployeeID       int     `json:"employee_id"`
	PaymentID        int     `json:"payment_id"`
	DeliveryID       int     `json:"delivery_id"`
	DateOrderID      int     `json:"date_order_id"`
	DiscountID       int     `json:"discount_id"`
	TotalAmountOrder float64 `json:"total_amount_order"`
	DiscountAmount   float64 `json:"discount_amount"`
	IsDelete         bool    `json:"is_delete"`
}
