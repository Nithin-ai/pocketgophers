package models

type Product struct {
	ProductID    string  `json:"id"`
	ProductName  string  `json:"prodname"`
	ProductDesc  string  `json:"proddesc"`
	ProductPrice float64 `json:"prodprice"`
	CreatedAt    string  `json:"createdAt"`
	Count        int     `json:"count"`
	CategoryName string  `json:"categoryname"`
}
