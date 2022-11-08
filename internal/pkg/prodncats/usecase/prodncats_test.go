package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mb_api/internal/pkg/models"
	"mb_api/internal/pkg/prodncats"
	"mb_api/internal/pkg/prodncats/mocks"
	"testing"
)

var (
	testProducts = models.Products{
		ProductName: "name",
		ProductCategory: "category",
	}
	testCategories = models.Categories{
		CategoryName: "name",
		CategoryProduct: "product",
	}
	testPairs = models.Pairs{
		ProductName: "pname",
		CategoryName: "cname",
	}
)

func getTestUseCase(mockRep *mocks.MockRepository) prodncats.UseCase {
	return &prodncatsUseCase{rep: mockRep}
}

func Test_prodncatsUseCase_InitCategories(t *testing.T) {
	// Create mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRep := mocks.NewMockRepository(mockCtrl)
	eUC := getTestUseCase(mockRep)
	var categories []models.Categories
	mockRep.EXPECT().GetCategories().Return(categories, nil)
	_, err := eUC.InitCategories(new([]models.Categories))

	assert.Equal(t, nil, err)
}

func Test_prodncatsUseCase_InitPairs(t *testing.T) {
	// Create mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRep := mocks.NewMockRepository(mockCtrl)
	eUC := getTestUseCase(mockRep)
	var pairs []models.Pairs
	mockRep.EXPECT().GetPairs().Return(pairs, nil)
	_, err := eUC.InitPairs(new([]models.Pairs))

	assert.Equal(t, nil, err)
}

func Test_prodncatsUseCase_InitProducts(t *testing.T) {
	// Create mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRep := mocks.NewMockRepository(mockCtrl)
	eUC := getTestUseCase(mockRep)
	var products []models.Products
	mockRep.EXPECT().GetProducts().Return(products, nil)
	_, err := eUC.InitProducts(new([]models.Products))

	assert.Equal(t, nil, err)
}