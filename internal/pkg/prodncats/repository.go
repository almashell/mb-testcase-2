package prodncats

import (
	"mb_api/internal/pkg/models"
)

type UseCase interface {
	InitProducts(products *[]models.Products) (int, error)
	InitCategories(categories *[]models.Categories) (int, error)
	InitPairs(pairs *[]models.Pairs) (int, error)
}
