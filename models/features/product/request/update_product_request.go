package request

type UpdateProductRequest struct {
	Id             int    `validate:"required"`
	ProductName    string `validate:"required;min=1;max=255" json:"productName"`
	AvailableStock int    `validate:"requeired"`
}
