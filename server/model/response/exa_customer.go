package response

import "quan/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
