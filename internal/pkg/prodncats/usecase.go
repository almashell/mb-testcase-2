package prodncats

import "mb_api/internal/pkg/models"

type Repository interface {
	GetProducts() ([]models.Products, error)
	GetCategories() ([]models.Categories, error)
	GetPairs() ([]models.Pairs, error)
}
