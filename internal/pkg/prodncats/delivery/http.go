package delivery

import (
	"mb_api/internal/pkg/models"
	"mb_api/internal/pkg/network"
	"mb_api/internal/pkg/prodncats"
	"mb_api/internal/pkg/prodncats/usecase"
	"net/http"
)

type prodncatsDelivery struct {
	UseCase prodncats.UseCase
}

func GetDelivery() prodncats.Delivery {
	return &prodncatsDelivery{
		UseCase: usecase.GetUseCase(),
	}
}

func (pd *prodncatsDelivery) GetProducts(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	var products []models.Products
	if code, err := pd.UseCase.InitProducts(&products); err != nil {
		network.GenErrorCode(w, r, err.Error(), code)
		return
	}

	network.Jsonify(w, products, http.StatusOK)
}

func (pd *prodncatsDelivery) GetCategories(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	var categories []models.Categories
	if code, err := pd.UseCase.InitCategories(&categories); err != nil {
		network.GenErrorCode(w, r, err.Error(), code)
		return
	}

	network.Jsonify(w, categories, http.StatusOK)
}

func (pd *prodncatsDelivery) GetPairs(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	var pairs []models.Pairs
	if code, err := pd.UseCase.InitPairs(&pairs); err != nil {
		network.GenErrorCode(w, r, err.Error(), code)
		return
	}

	network.Jsonify(w, pairs, http.StatusOK)
}