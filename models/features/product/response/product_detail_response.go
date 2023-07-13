package response

type ProductDetailResponse struct {
	Id             int    `json:"id"`
	ProductName    string `json:"productName"`
	AvailableStock int    `json:"availableStock"`
}