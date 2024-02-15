package campaign

import (
	"fmt"

	"mono.thienhang.com/pkg/models/base"
)

type Campaign struct {
	base.Base
	Name        *string `json:"name"`
	TimeStart   *string `json:"time_start"`
	TimeEnd     *string `json:"time_end"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
}

func (c *Campaign) Validate() []error {
	errs := make([]error, 0, 5)
	// Call the base validation
	// if err := c.Base.Validate(); err != nil {
	// 	return err
	// }

	// Additional Campaign-specific validations
	if c.Name == nil {
		err := fmt.Errorf("Name cannot be empty")
		errs = append(errs, err)
	}

	// You can add more validation rules as needed

	return errs
}
