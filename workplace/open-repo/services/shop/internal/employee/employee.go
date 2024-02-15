package employee

import "mono.thienhang.com/pkg/models/base"

type Employee struct {
	base.Base
	Name         string `json:"name"`
	City         string `json:"city"`
	Country      string `json:"country"`
	StreetName   string `json:"street_name"`
	StreetNumber string `json:"street_number"`
	Gender       string `json:"gender"`
	Age          int    `json:"age"`
	SalaryLevel  string `json:"salary_level"`
	Status       string `json:"status"`
	// StartWorkDay   string  `json:"start_work_day"`
	EducationLevel string  `json:"education_level"`
	LastRating     float64 `json:"last_rating"`
}
