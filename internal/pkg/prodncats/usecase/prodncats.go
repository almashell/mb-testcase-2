package usecase

//go:generate mockgen -destination=../mocks/mock_usecase.go -package=mocks mb_api/internal/pkg/prodncats UseCase

import (
	"mb_api/internal/pkg/db"
	"mb_api/internal/pkg/models"
	"mb_api/internal/pkg/prodncats"
	"mb_api/internal/pkg/prodncats/repository"
	"net/http"
)

type prodncatsUseCase struct {
	rep prodncats.Repository
}

func GetUseCase() prodncats.UseCase {
	return &prodncatsUseCase{
			rep: repository.NewSqlProdncatsRepository(db.ConnectToDB()),
	}
}

func (ec *prodncatsUseCase) InitProducts(products *[]models.Products) (status int, err error) {
	*products, err = ec.rep.GetProducts()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (ec *prodncatsUseCase) InitCategories(categories *[]models.Categories) (status int, err error) {
	*categories, err = ec.rep.GetCategories()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (ec *prodncatsUseCase) InitPairs(pairs *[]models.Pairs) (status int, err error) {
	*pairs, err = ec.rep.GetPairs()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}