package models

import "net/http"

type WorkMessage struct {
	Request *http.Request `json:"-"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
}

type Products struct {
	ProductName     string `json:"prodname"`
	ProductCategory string `json:"categories, omitempty"`
}

type Categories struct {
	CategoryName    string `json:"catname"`
	CategoryProduct string `json:"products, omitempty"`
}

type Pairs struct {
	ProductName  string `json:"prodname"`
	CategoryName string `json:"catname, omitempty"`
}
