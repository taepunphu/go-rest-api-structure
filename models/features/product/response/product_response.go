package response

type ProductResponse struct {
	Id             int    `json:"id"`
	ProductName    string `json:"productName"`
	AvailableStock int    `json:"availableStock"`
}
