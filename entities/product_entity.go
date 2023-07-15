package entities

type ProductEntity struct {
	Id             int    `gorm:"column:id; type:int; primaryKey"`
	ProductName    string `gorm:"column:product_name; type:varchar(255)"`
	AvailableStock int    `gorm:"column:available_stock; type:int"`
}
