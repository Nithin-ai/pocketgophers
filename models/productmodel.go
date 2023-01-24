package models

type Product struct {
	ProductID    uint    `json:"id"`
	ProductName  string  `json:"prodname"`
	ProductDesc  string  `json:"proddesc"`
	ProductPrice float64 `json:"prodprice"`
	CreatedAt    string  `json:"createdAt"`
	Count        int     `json:"count"`
	CategoryName string  `json:"categoryname"`
}
