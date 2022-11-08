package prodncats

import "net/http"

type Delivery interface {
	GetProducts(w http.ResponseWriter, r *http.Request, ps map[string]string)
	GetCategories(w http.ResponseWriter, r *http.Request, ps map[string]string)
	GetPairs(w http.ResponseWriter, r *http.Request, ps map[string]string)
}