package web

type AddressUpsertRequest struct {
	Street     string `json:"street" validate:"required,min=1,max=255"`
	City       string `json:"city" validate:"required,min=1,max=255"`
	Province   string `json:"province" validate:"required,min=1,max=255"`
	Country    string `json:"country" validate:"required,min=1,max=255"`
	PostalCode string `json:"postal_code" validate:"required,min=1,max=255"`
}