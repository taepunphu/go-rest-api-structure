package request

type CreateProductRequest struct {
	ProductName    string `validate:"required, min=1, max=255" json:"productName"`
	AvailableStock int    `validate:"required" json:"availableStock"`
}
